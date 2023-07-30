package interfaces

import (
	"github.com/abhinandkakkadi/moviesgo-admin-service/pkg/utils/models"
)

type AdminUseCase interface {
	AdminSignUp(user models.AdminSignUp) (int, error)
	AdminLogin(userSignUpDetails models.AdminLogin) (string, error)
	ValidateAdmin(token string) error
}
