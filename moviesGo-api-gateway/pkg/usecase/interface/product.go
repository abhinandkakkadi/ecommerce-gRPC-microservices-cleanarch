package interfaces

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type ProductUseCase interface {
	ShowAllProducts(page int, count int) ([]models.ProductOfferBriefResponse, error)
	ShowAllProductsToAdmin(page int, count int) ([]models.ProductsBrief, error)
	ShowIndividualProducts(sku string) (models.ProductOfferLongResponse, error)
	AddProduct(product models.ProductsReceiver) (models.ProductResponse, error)
}
