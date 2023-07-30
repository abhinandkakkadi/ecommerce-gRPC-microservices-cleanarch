package usecase

import (
	"errors"

	"github.com/abhinandkakkadi/moviesgo-admin-service/pkg/helper"
	interfaces "github.com/abhinandkakkadi/moviesgo-admin-service/pkg/repository/interface"
	services "github.com/abhinandkakkadi/moviesgo-admin-service/pkg/usecase/interface"
	"github.com/abhinandkakkadi/moviesgo-admin-service/pkg/utils/models"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repo,
	}
}


func (a *adminUseCase) ValidateAdmin(token string) error {

	err := helper.ValidateAdminToken(token)
	if err != nil {
		return err
	}

	return nil
	
}

func (a *adminUseCase) AdminSignUp(adminDetails models.AdminSignUp) (int, error) {
	// validator package to check the constraints specified in the struct which is used to retrieve these details
	if err := validator.New().Struct(adminDetails); err != nil {
		return 400, err
	}

	// check whether the admin already exist in the database -
	userExist := a.adminRepository.CheckAdminAvailability(adminDetails)
	if userExist {
		return 400, errors.New("admin already exist, sign in")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminDetails.Password), 10)
	if err != nil {
		return 500, errors.New("internal server error")
	}
	adminDetails.Password = string(hashedPassword)

	status, err := a.adminRepository.CreateAdmin(adminDetails)
	if err != nil {
		return 400, err
	}

	return status, nil

}

func (a *adminUseCase) AdminLogin(adminDetails models.AdminLogin) (string, error) {

	// getting details of the admin based on the email provided
	adminCompareDetails, err := a.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return "", err
	}

	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return "", err
	}

	
	token, err := helper.GenerateTokenAdmin(adminCompareDetails)

	if err != nil {
		return "", err
	}	

	return token,nil
}
