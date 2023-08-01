package interfaces

import (
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"
	"github.com/gin-gonic/gin"
)

type UserClient interface {
	SampleRequest(request string) (string, error)
	SignUpRequest(userDetails models.UserDetails) (int, error)
	LoginHandler(userAuthDetails models.UserLogin) (string, error)
	AddAddress(address models.AddressInfo, userID int) (int, error)
	GetAllAddress(userID int) ([]models.AddressInfoResponse,error)
	Checkout(userID int) (models.CheckoutDetails,error)
	UserAuthRequired(c *gin.Context)
}
