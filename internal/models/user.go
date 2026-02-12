package models

import "time"

type Users struct {
	ID        uint  `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `gorm:"uniqueIndex" json:"email"`
	Password  string `json:"password"`
	Role      string `gorm:"default:user" json:"role"`
	RefreshToken string `json:"-"`
	CreatedAt time.Time
	UpdatedAT time.Time
}