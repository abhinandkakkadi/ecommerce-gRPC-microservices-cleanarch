package interfaces

import "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"


type UserRepository interface {
	UserSignUp(user models.UserDetails) (int, error)
}
