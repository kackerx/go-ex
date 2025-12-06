package main

import "context"

type PubSuber interface {
	Publish(ctx context.Context, topic string, message any) error
	Subscribe(ctx context.Context, topic string) (<-chan any, error)
}
