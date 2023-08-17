package interfaces

type ProductClient interface {
	DoesProductExist(productID int) (bool, error)
	GetProductPriceFromID(productID int) (float64, error)
	GetProductNamesFromID(productId []int) (map[int]string, error)
}
