package interfaces

import "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"

type CartClient interface {
	AddToCart(productID int,userId int) (int,error)
	DisplayCart(userID int) (models.CartResponse,error)
}
