package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
