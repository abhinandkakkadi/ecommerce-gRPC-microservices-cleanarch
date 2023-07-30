package repository

import (
	interfaces "github.com/abhinandkakkadi/moviesgo-products-service/pkg/repository/interface"
	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/utils/models"
	"gorm.io/gorm"
)

type ProductDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductsRepository {
	return &ProductDatabase{
		DB: DB,
	}
}


func (p *ProductDatabase) ShowAllProducts(page int, count int) (models.ProductsBrief, error) {

}

func (p *ProductDatabase) AddProduct() {

}

func (p *ProductDatabase) GetGenreDetails() {

} 

func (p *ProductDatabase) GetStudioDetails() {

}



