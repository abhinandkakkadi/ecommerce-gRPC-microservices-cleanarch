package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID     uint     `json:"user_id" gorm:"uniquekey; not null"`
	ProductID  uint     `json:"product_id"`
	Quantity   float64  `json:"quantity"`
	TotalPrice float64  `json:"total_price"`
}