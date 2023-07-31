package interfaces

import (
	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/utils/models"
)

type ProductsRepository interface {
	ShowAllProducts(page int, count int) ([]models.ProductsBrief, error)
	AddProduct(productReceiver models.ProductsReceiver) (int, error)
	GetGenreDetails()
	GetStudioDetails()
	DoesProductExist(productID int) (bool, error)
	GetProductPriceFromID(productID int) (float64, error)
	GetProductNameFromID(productIDS []int) ([]models.ProductName, error)
}
