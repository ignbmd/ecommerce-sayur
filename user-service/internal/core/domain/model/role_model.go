package model

import "time"

type Role struct {
	Id        int64
	Name      string
	Users     []User `gorm:"many2many:user_role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
