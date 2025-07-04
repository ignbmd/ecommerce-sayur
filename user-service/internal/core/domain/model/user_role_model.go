package model

import "time"

type UserRole struct {
	ID        int64 `gorm:"primaryKey"`
	UserId    int64 `gorm:"index"`
	RoleId    int64 `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Tabler interface {
	TableName() string
}

func (UserRole) TableName() string {
	return "user_role"
}
