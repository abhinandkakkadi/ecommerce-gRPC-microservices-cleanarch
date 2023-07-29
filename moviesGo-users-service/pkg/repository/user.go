package repository

import (
	"github.com/abhinandkakkadi/moviesgo-user-service/pkg/repository/interfaces"
	"github.com/abhinandkakkadi/moviesgo-user-service/pkg/utils/models"
	"gorm.io/gorm"
)



type UserDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &UserDatabase{DB}
}


func (u *UserDatabase) UserSignUp(user models.UserDetails) (int, error) {
	
}