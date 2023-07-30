package interfaces

import (
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"
	"github.com/gin-gonic/gin"
)


type AdminClient interface {

	LoginHandler(adminDetails models.AdminLogin) (string,error)
	CreateAdmin(admin models.AdminSignUp) (int,error)
	AdminAuthRequired(c *gin.Context)
	
}