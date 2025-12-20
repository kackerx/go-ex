package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type StandardHandler interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func main() {

	r := gin.Default()
	v1 := r.Group("/v1")

	userRepo := NewUserRepo()
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	RegisterRouter(
		v1,
		"users",
		userHandler,
		WithCustomMethod("PATCH", "enable", userHandler.Enable),
		WithCustomMethod("PATCH", "disable", userHandler.Disable),
	)

	slog.Info("starting server on port 8080")
	r.Run(":8080")
}
