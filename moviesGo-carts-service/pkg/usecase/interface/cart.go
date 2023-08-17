package interfaces

import (
	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/domain"
	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/utils/models"
)

type CartUseCase interface {
	AddToCart(productID int, userID int) (int, error)
	DisplayCArt(userID int) (models.CartResponse, error)
	OrderItemsFromCarts(orderItems domain.Order) (int, error)
}
