package interfaces

import (
	"github.com/abhinandkakkadi/moviesgo-admin-service/pkg/domain"
	"github.com/abhinandkakkadi/moviesgo-admin-service/pkg/utils/models"
)

type AdminRepository interface {
	CreateAdmin(admin models.AdminSignUp) (int, error)
	CheckAdminAvailability(admin models.AdminSignUp) bool
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
}
