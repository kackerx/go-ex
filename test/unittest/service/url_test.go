package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Package Setup")

	code := m.Run()

	fmt.Println("Package Teardown")

	os.Exit(code)
}

func TestIsValidURL(t *testing.T) {
	if os.Getenv("env") == "local" {
		t.Skip("skip test in local environment")
	}

	fmt.Println("TestIsValidURL Setup")
	t.Cleanup(func() {
		fmt.Println("TestIsValidURL Teardown")
	})

	testCase := []struct {
		name   string
		rawURL string

		before func(t *testing.T)
		after  func(t *testing.T)

		want bool
	}{
		{
			name:   "valid https url",
			rawURL: "https://www.google.com",
			before: func(t *testing.T) {
				fmt.Println("TestValidURL Setup")
			},
			after: func(t *testing.T) {
				fmt.Println("TestValidURL Teardown")
			},
			want: true,
		},
		{
			name:   "valid http url",
			rawURL: "http://www.google.com",
			want:   true,
		},
		{
			name:   "missing scheme",
			rawURL: "www.google.com",
			want:   false,
		},
		{
			name:   "missing host",
			rawURL: "https://",
			want:   false,
		},
		{
			name:   "invalid scheme",
			rawURL: "ftp://www.google.com",
			want:   false,
		},
		{
			name:   "null string",
			rawURL: "",
			want:   false,
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before(t)
			}

			t.Cleanup(func() {
				if tt.after != nil {
					tt.after(t)
				}
			})

			got := IsValidURL(tt.rawURL)
			assert.Equal(t, tt.want, got, "rawURL: %s", tt.rawURL)
		})
	}
}
