package models

import "time"

type OrderItem struct {
	ID uint `gorm:"primaryKey" json:"id"`
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}