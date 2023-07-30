package usecase

import (
	interfaces "github.com/abhinandkakkadi/moviesgo-products-service/pkg/repository/interface"
	services "github.com/abhinandkakkadi/moviesgo-products-service/pkg/usecase/interface"
	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/utils/models"
)

type adminUseCase struct {
	adminRepository interfaces.ProductsRepository
}

func NewProductUseCase(repo interfaces.ProductsRepository) services.ProductUseCase {
	return &adminUseCase{
		adminRepository: repo,
	}
}


func (a *adminUseCase) ShowAllProducts(page int, count int) (models.ProductsBrief, error) {

}

func (a *adminUseCase) 	AddProduct()  {
	
}

func (a *adminUseCase) GetGenreDetails() {
	
}

func (a *adminUseCase) GetStudioDetails() {
	
}




 
	
