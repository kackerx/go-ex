package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kackerx/go-ex/test/unittest/domain"
	"github.com/kackerx/go-ex/test/unittest/service"
)

func NewLinkHandler(service service.LinkService) *LinkHandler {
	return &LinkHandler{
		service: service,
	}
}

type LinkHandler struct {
	service service.LinkService
}

type CreateLinkReq struct {
	URL string `json:"url"`
}

type CreateLinkResp struct {
	Code string `json:"code"`
}

func (h *LinkHandler) CreateLink(ctx *gin.Context) {
	req, err := ParseReq[CreateLinkReq](ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link, err := h.service.CreateLink(ctx.Request.Context(), &domain.Link{
		OriginalURL: req.URL,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, CreateLinkResp{
		Code: link.ShortCode,
	})
}

func ParseReq[T any](ctx *gin.Context) (T, error) {
	var req T
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return req, err
	}
	return req, nil
}
