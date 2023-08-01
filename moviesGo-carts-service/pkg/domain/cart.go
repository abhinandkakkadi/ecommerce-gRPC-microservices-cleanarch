package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID     uint     `json:"user_id" gorm:"uniquekey; not null"`
	ProductID  uint     `json:"product_id"`
	Quantity   float64  `json:"quantity"`
	TotalPrice float64  `json:"total_price"`
}

type Order struct {
	ID         uint `gorm:"primaryKey"`
	UserID     int
	AddressID  int
	TotalPrice float64
	Items      []OrderItems `gorm:"foreignKey:OrderID"` 
}

type OrderItems struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID int
	Quantity  int
	Price     float64
}