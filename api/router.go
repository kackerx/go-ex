package main

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup, resourceName string, ctrl StandardHandler, opts ...Option) {
	group := r.Group("/" + resourceName)
	{
		group.GET("", ctrl.List)
		group.GET("/:id", ctrl.Get)
		group.POST("", ctrl.Create)
		group.PATCH("/:id", ctrl.Update)
		group.DELETE("/:id", ctrl.Delete)
	}
	for _, opt := range opts {
		opt(group)
	}
}

type Option func(group *gin.RouterGroup)

func WithCustomMethod(method string, action string, handler func(ctx *gin.Context)) Option {
	return func(group *gin.RouterGroup) {
		group.Handle(method, "/:id/"+action, handler)
	}
}
