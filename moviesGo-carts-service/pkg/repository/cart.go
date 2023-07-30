package repository

import (
	interfaces "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/repository/interface"
	"gorm.io/gorm"
)

type CartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &CartDatabase{
		DB: DB,
	}
}

