package interfaces

import (
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"
	"github.com/gin-gonic/gin"
)

type UserClient interface {
	SampleRequest(request string) (string, error)
	SignUpRequest(userDetails models.UserDetails) (int, error)
	LoginHandler(userAuthDetails models.UserLogin) (string,error)
	UserAuthRequired(c *gin.Context)
}
