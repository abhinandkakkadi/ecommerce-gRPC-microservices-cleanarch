package usecase

import (
	"errors"

	"github.com/abhinandkakkadi/moviesgo-user-service/pkg/helper"
	interfaces "github.com/abhinandkakkadi/moviesgo-user-service/pkg/repository/interface"
	services "github.com/abhinandkakkadi/moviesgo-user-service/pkg/usecase/interface"
	"github.com/abhinandkakkadi/moviesgo-user-service/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (u *userUseCase) UserSignUp(userDetails models.UserDetails) (int,error) {

		// Check whether the user already exist. If yes, show the error message, since this is signUp
		userExist := u.userRepo.CheckUserAvailability(userDetails.Email)

		if userExist {
			return 400, errors.New("user already exist, sign in")
		}
	
		if userDetails.Password != userDetails.ConfirmPassword {
			return 400, errors.New("password does not match")
		}
	
		// Hash password since details are validated
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDetails.Password), 10)
		if err != nil {
			return 500, errors.New("internal server error")
		}
		
		userDetails.Password = string(hashedPassword)
	
		// add user details to the database
		_, err = u.userRepo.AddUserDetails(userDetails)
		if err != nil {
			return 500, err
		}

		return 201,nil
	
}

func (u *userUseCase) UserLogin(userDetails models.UserLogin) (string,error) {

	// checking if a username exist with this email address
	ok := u.userRepo.CheckUserAvailability(userDetails.Email)
	if !ok {
		return "", errors.New("the user does not exist")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user, err := u.userRepo.FindUserByEmail(userDetails)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password))
	if err != nil {
		return "", errors.New("password incorrect")
	}

	token, err := helper.GenerateAccessToken(user)
	if err != nil {
		return "", errors.New("could not create token")
	}

	return token,nil

}

	
	

