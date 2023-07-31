package interfaces

import (
	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/utils/models"
)

type ProductUseCase interface {
	ShowAllProducts(page int, count int) ([]models.ProductsBrief, error)
	AddProduct(productReceiver models.ProductsReceiver) (int, error)
	GetGenreDetails()
	GetStudioDetails()
	ProductExistInCarts(productID int) (bool, error)
	GetProductPriceFromID(productID int) (float64, error)
	GetProductsNameFromID(productIDS []int) ([]models.ProductName, error)
}
