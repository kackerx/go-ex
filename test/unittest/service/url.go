package service

import (
	"context"
	"net/url"
	"slices"

	"github.com/kackerx/go-ex/test/unittest/domain"
)

type LinkService interface {
	CreateLink(ctx context.Context, link *domain.Link) (*domain.Link, error)
}

func IsValidURL(rawURL string) bool {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	if u.Host == "" {
		return false
	}

	return slices.Contains([]string{"http", "https"}, u.Scheme)
}
