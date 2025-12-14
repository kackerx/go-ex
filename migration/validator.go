package main

import (
	"context"
	"log/slog"

	"gorm.io/gorm"
)

type Validator[T Entity] struct {
	base     *gorm.DB // 第二阶段: base是源表, 第三阶段: base是目标表
	target   *gorm.DB
	producer Producer

	direction Direction
	// orderFields []string
}

func (v *Validator[T]) Validate(ctx context.Context) error {
	var (
		base   T
		offset = -1
	)
	for {
		offset++
		err := v.base.Offset(offset).Order("id").First(&base).Error
		switch err {
		case gorm.ErrRecordNotFound:
			return nil
		case nil:
			if err := v.Compare(ctx, base); err != nil {
				slog.Error("compare error", "base_id", base.ID(), "error", err)
				continue
			}
		default:
			slog.Error("db error", "error", err)
			continue
		}
	}
}

func (v *Validator[T]) Compare(ctx context.Context, base T) error {
	var (
		target           T
		inconsistentType InconsistentType
	)
	err := v.target.Where("id = ?", base.ID()).First(&target).Error
	switch err {
	case gorm.ErrRecordNotFound:
		inconsistentType = InconsistentTypeTargetMissing
	case nil:
		if !base.Equal(target) {
			inconsistentType = InconsistentTypeNotEqual
		}
	default:
		return err
	}

	if inconsistentType == "" {
		return nil
	}

	return v.producer.PublishInconsistentEvent(ctx, &InconsistentEvent{
		ID:        base.ID(),
		Direction: v.direction,
		Type:      inconsistentType,
	})
}
