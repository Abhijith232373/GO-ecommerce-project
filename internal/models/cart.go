package models

import "time"

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`

	Product Product `gorm:"foreginKey:ProductID" json:"product"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}