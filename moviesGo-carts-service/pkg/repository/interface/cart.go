package interfaces

import (
	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/domain"
	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/utils/models"
)

type CartRepository interface {
	DoesProductExistInCart(productID int) (bool,error)
	InsertNewProductInCart(productID int,quantity int,productPrice float64,userID int) error
	IterateQuantityInCart(productID int,userID int,productPrice float64) error
	GetUserCartFromID(userID int) ([]models.Cart,error)
	CreateNewOrder(orderProducts domain.Order) (int,error)
	DeleteCartItems(userID int) error

}
