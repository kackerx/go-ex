package po

import "time"

type BaseModel struct {
	ID        int    `gorm:"primary_key"`
	CreatedBy string `gorm:"column:created_by;type:varchar(20)"`
	UpdateBy  string `gorm:"column:update_by;type:varchar(20)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDeleted bool
}
