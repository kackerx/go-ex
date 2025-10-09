package main

import (
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"go-bobo/adapter/handlers/rpc"
	"go-bobo/app/executor/user"
	"go-bobo/infra/persistence/dal"
	"go-bobo/infra/persistence/repository"
	"go-bobo/pkg/config"
	"go-bobo/server"

	userPB "go-bobo/api/user"
)

func main() {
	runApp()
}

func runApp() {
	cnf := config.NewConfig()
	db := newDB(cnf)
	userDal := dal.NewUserDal(db)
	userRepo := repository.NewUserRepositoryImpl(userDal)
	userApp := user.NewUserApplication(userRepo)

	server.RunGRPCServer(func(server *grpc.Server) {
		// 注册服务handler
		userHandler := rpc.NewUserHandler(userApp)
		userPB.RegisterUserServer(server, userHandler)
	})
}

func newDB(cnf *viper.Viper) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cnf.Get("db.username"),
		cnf.Get("db.password"),
		cnf.Get("db.host"),
		cnf.Get("db.port"),
		cnf.Get("db.db_name"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}
