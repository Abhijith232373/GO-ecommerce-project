package models

import "time"

type Order struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	UserID      uint    `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `gorm:"default:'pending'" json:"status"`

	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	City     string `json:"city"`
	State    string `json:"state"`
	Pincode  string `json:"pincode"`

	OrderItems []OrderItem `gorm:"foreginKey:OrderID" json:"items"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}