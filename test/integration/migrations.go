package integration

import (
	"embed"

	"github.com/uptrace/bun/migrate"
)

//go:embed *.sql
var sqlMigrations embed.FS

// Migrations holds all database migrations
var Migrations = migrate.NewMigrations()

func init() {
	// Register SQL migrations
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}
