package models

import (
	"time"
	"gorm.io/gorm"
)

type Users struct {
	ID        uint  `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `gorm:"uniqueIndex" json:"email"`
	Password  string `json:"password"`
	Role      string `gorm:"default:user" json:"role"`
	RefreshToken string `json:"-"`
	IsActive bool
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}