package models

import "time"

type Wishlist struct {
	ID uint `gorm:"primaryKey"`

	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`

	Product Product `gorm:"foreignKey:ProductID" json:"product"`

	CreatedAt time.Time `json:"created_at"`
}