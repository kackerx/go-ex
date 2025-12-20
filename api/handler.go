package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CRUDService[T, CreateReq, UpdateReq any] interface {
	Create(ctx context.Context, req *CreateReq) (*T, error)
	Update(ctx context.Context, req *UpdateReq) (*T, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*T, error)
	List(ctx context.Context, req ListRequest) (*ListResponse[T], error)
}

type BaseHandler[T, CreateReq, UpdateReq any] struct {
	service CRUDService[T, CreateReq, UpdateReq]
}

func NewBaseHandler[T, CreateReq, UpdateReq any](service CRUDService[T, CreateReq, UpdateReq]) *BaseHandler[T, CreateReq, UpdateReq] {
	return &BaseHandler[T, CreateReq, UpdateReq]{
		service: service,
	}
}

func (h *BaseHandler[T, C, U]) Create(ctx *gin.Context) {
	var req C
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (h *BaseHandler[T, C, U]) Update(ctx *gin.Context) {
	var req U
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.Update(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (h *BaseHandler[T, C, U]) Delete(ctx *gin.Context) {
	err := h.service.Delete(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *BaseHandler[T, C, U]) Get(ctx *gin.Context) {
	item, err := h.service.Get(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (h *BaseHandler[T, C, U]) List(ctx *gin.Context) {
	var req ListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, err := h.service.List(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, list)
}
