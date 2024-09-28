package main

import (
	"context"
	"fmt"
)

type Handler func(ctx context.Context)

type Middleware func(handler Handler) Handler

type Chain struct {
	mids []Middleware
}

func NewChain() *Chain {
	return &Chain{
		mids: make([]Middleware, 0),
	}
}

func (c *Chain) AddMid(mid Middleware) {
	c.mids = append(c.mids, mid)
}

func (c *Chain) Do(ctx context.Context) {
	root := func(ctx context.Context) {
		fmt.Println("i am root")
	}

	for i := len(c.mids) - 1; i >= 0; i-- {
		root = c.mids[i](root)
	}

	root(ctx)
}

func main() {
	c := NewChain()
	c.AddMid(func(handler Handler) Handler {
		return func(ctx context.Context) {
			fmt.Println("1111111-before")
			handler(ctx)
			fmt.Println("1111111-after")
		}
	})

	c.AddMid(func(handler Handler) Handler {
		return func(ctx context.Context) {
			fmt.Println("2222222-before")
			handler(ctx)
			fmt.Println("2222222-after")
		}
	})

	c.AddMid(func(handler Handler) Handler {
		return func(ctx context.Context) {
			fmt.Println("3333333-before")
			handler(ctx)
			fmt.Println("3333333-after")
		}
	})

	c.Do(context.Background())
}
