package interfaces

type ProductClient interface {
	DoesProductExist(productID int) (bool,error)
}
