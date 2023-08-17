package domain

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
}

type Address struct {
	gorm.Model
	UserID    uint
	Users     Users `gorm:"foreignkey:UserID"`
	Name      string
	HouseName string
	City      string
}

type AddressResponse struct {
	ID        int
	Name      string
	HouseName string
	City      string
}
