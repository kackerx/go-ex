package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*BaseHandler[User, UserCreateReq, UserUpdateReq]
	userService *UserService
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{
		BaseHandler: NewBaseHandler(userService),
		userService: userService,
	}
}

func (h *UserHandler) Enable(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "enable"})
}

func (h *UserHandler) Disable(ctx *gin.Context) {
}
