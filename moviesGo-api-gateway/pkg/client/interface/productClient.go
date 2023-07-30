package interfaces

import "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"

type ProductClient interface {
	ShowAllProducts(page int, count int) ([]models.ProductsBrief, error)
	AddProduct(productDetails models.ProductsReceiver) (int, error)
	// GetGenreDetails()
	// GetStudioDetails()
}
