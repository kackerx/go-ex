package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kackerx/go-ex/test/unittest/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type StubLinkService struct {
	CreateLinkFunc func(ctx context.Context, link *domain.Link) (*domain.Link, error)
}

// CreateLink implements service.LinkService.
func (s *StubLinkService) CreateLink(ctx context.Context, link *domain.Link) (*domain.Link, error) {
	return s.CreateLinkFunc(ctx, link)
}

func TestLinkHandler_CreateLink(t *testing.T) {
	testCase := []struct {
		name string
		stub *StubLinkService
		req  CreateLinkReq
		want CreateLinkResp
	}{
		{
			name: "success",
			stub: &StubLinkService{
				CreateLinkFunc: func(ctx context.Context, link *domain.Link) (*domain.Link, error) {
					return &domain.Link{
						ShortCode: "123456",
					}, nil
				},
			},
			req: CreateLinkReq{
				URL: "https://www.google.com",
			},
			want: CreateLinkResp{
				Code: "123456",
			},
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewLinkHandler(tt.stub)

			req := httptest.NewRequest("POST", "/links", bytes.NewBufferString(`{"url": "https://www.google.com"}`))
			router := gin.Default()
			router.POST("/links", handler.CreateLink)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			var got CreateLinkResp
			err := json.Unmarshal(w.Body.Bytes(), &got)
			require.NoError(t, err)

			assert.Equal(t, tt.want, got)
		})
	}

}
