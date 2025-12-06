package integration

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// ArticleTestSuite 测试套件
type ArticleTestSuite struct {
	suite.Suite
	server *gin.Engine
}

func (s *ArticleTestSuite) TestEdit() {
	t := s.T()
	testCase := []struct {
		name string

		before func(t *testing.T)
		after  func(t *testing.T)

		req *Article

		wantCode   int
		wantResult *Result[int64]
	}{
		{
			name: "test abc",
			req: &Article{
				Title:   "test",
				Content: "test",
			},

			wantCode: 200,
			wantResult: &Result[int64]{
				Code:    0,
				Message: "success",
			},
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {

		})
	}

}

func TestArticle(t *testing.T) {
	suite.Run(t, new(ArticleTestSuite))
}

type Article struct {
	Title   string
	Content string
}

type Result[T any] struct {
	Code    int
	Message string
	Data    T
}
