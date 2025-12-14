package main

import (
	"context"

	"gorm.io/gorm"
)

type Fixer struct {
	base     *gorm.DB
	target   *gorm.DB
	producer Producer

	direction Direction
	// orderFields []string
}

func (f *Fixer) Fix(ctx context.Context, event *InconsistentEvent) error {
	switch event.Type {
	case InconsistentTypeTargetMissing: // insert
		return nil
	case InconsistentTypeNotEqual: // update
		return nil
	case InconsistentTypeBaseMissing: // delete
		return nil
	default:
		return nil
	}
}
