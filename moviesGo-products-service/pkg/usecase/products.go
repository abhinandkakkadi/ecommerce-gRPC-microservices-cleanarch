package usecase

import (
	"errors"

	interfaces "github.com/abhinandkakkadi/moviesgo-products-service/pkg/repository/interface"
	services "github.com/abhinandkakkadi/moviesgo-products-service/pkg/usecase/interface"
	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/utils/models"
)

type productUseCase struct {
	productRepository interfaces.ProductsRepository
}

func NewProductUseCase(repo interfaces.ProductsRepository) services.ProductUseCase {

	return &productUseCase{
		productRepository: repo,
	}
}


func (p *productUseCase) ShowAllProducts(page int, count int) ([]models.ProductsBrief, error) {

	return p.productRepository.ShowAllProducts(page,count)

}

func (p *productUseCase) 	AddProduct(productReceiver models.ProductsReceiver) (int,error)  {

	id,err :=  p.productRepository.AddProduct(productReceiver)
	
	if err != nil {
		return 0,err
	}

	if id == 0 {
		return 0,errors.New("product not created")
	}

	return id,nil

}	

func (p *productUseCase) GetGenreDetails() {
	
}


func (p *productUseCase) GetStudioDetails() {
	
}




 
	
