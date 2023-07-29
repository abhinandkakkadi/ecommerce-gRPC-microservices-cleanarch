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
	Id        uint   `json:"id" gorm:"unique;not null"`
	UserID    uint   `json:"user_id"`
	Users     Users  `json:"-" gorm:"foreignkey:UserID"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}
