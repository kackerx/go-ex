package integration

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

type User struct {
	bun.BaseModel `bun:"users"`

	ID   int64  `bun:"id,pk,autoincrement"`
	Name string `bun:"name"`
	Age  int    `bun:"age"`
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserRepository struct {
	db *bun.DB
}

func (r *UserRepository) Create(ctx context.Context, user *User) (int64, error) {
	_, err := r.db.NewInsert().Model(user).Returning("id").Exec(ctx)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *UserRepository) Get(ctx context.Context, id int64) (*User, error) {
	user := &User{}
	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

var (
	db  *bun.DB
	ctx = context.Background()
)

func TestMain(m *testing.M) {
	req := testcontainers.ContainerRequest{
		Image: "postgres:18-alpine3.22",
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "testdb",
		},
		ExposedPorts: []string{"5432"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(60 * time.Second),
	}

	pgContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	defer pgContainer.Terminate(ctx)

	pgHost, err := pgContainer.Host(ctx)
	if err != nil {
		panic(err)
	}

	pgPort, err := pgContainer.MappedPort(ctx, "5432")
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", "test", "postgres", pgHost, pgPort.Port(), "testdb")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())
	defer db.Close()

	migrator := migrate.NewMigrator(db, Migrations)

	// Create migrations table if it doesn't exist
	if err := migrator.Init(context.Background()); err != nil {
		panic(err)
	}

	// Run pending migrations
	group, err := migrator.Migrate(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Migrations completed:", group.ID)

	code := m.Run()
	os.Exit(code)
}

func TestSaveAndFindUser(t *testing.T) {
	asserts := assert.New(t)
	requires := require.New(t)
	testCases := []struct {
		name string
		user *User
	}{
		{
			name: "save and find user",
			user: &User{
				Name: "John",
				Age:  20,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := NewUserRepository(db)
			id, err := userRepo.Create(ctx, tt.user)
			asserts.NoError(err)

			t.Log(id)

			user, err := userRepo.Get(ctx, id)
			requires.NoError(err)

			asserts.Equal(tt.user.Name, user.Name)
			asserts.Equal(tt.user.Age, user.Age)
		})
	}
}
