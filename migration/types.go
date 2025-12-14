package main

import "context"

type Entity interface {
	ID() int64
	Equal(other Entity) bool
}

type Direction string

const (
	DirectionSource      Direction = "source"
	DirectionDestination Direction = "destination"
)

type Producer interface {
	PublishInconsistentEvent(ctx context.Context, event *InconsistentEvent) error
}

type InconsistentEvent struct {
	ID        int64
	Direction Direction
	Type      InconsistentType
}

type InconsistentType string

const (
	InconsistentTypeBaseMissing   InconsistentType = "base_missing"
	InconsistentTypeTargetMissing InconsistentType = "target_missing"
	InconsistentTypeNotEqual      InconsistentType = "not_equal"
)
