package service

import (
	"context"
	"net/http"

	"github.com/abhinandkakkadi/moviesgo-admin-service/pkg/admin/pb"
	interfaces "github.com/abhinandkakkadi/moviesgo-admin-service/pkg/usecase/interface"
	"github.com/abhinandkakkadi/moviesgo-admin-service/pkg/utils/models"
)

type AdminServiceServer struct {
	adminUseCase interfaces.AdminUseCase
	pb.UnimplementedAdminAuthServiceServer
}

func NewAdminServiceServer(useCase interfaces.AdminUseCase) pb.AdminAuthServiceServer {

	return &AdminServiceServer{
		adminUseCase: useCase,
	}

}

func (a *AdminServiceServer) AdminSignUp(ctx context.Context, userSignUpDetails *pb.AdminSingUpRequest) (*pb.AdminSignUpResponse, error) {

	adminCreateDetails := models.AdminSignUp{
		Name:     userSignUpDetails.Name,
		Email:    userSignUpDetails.Email,
		Password: userSignUpDetails.Password,
	}

	status, err := a.adminUseCase.AdminSignUp(adminCreateDetails)
	if err != nil {
		return &pb.AdminSignUpResponse{
			Status: int64(status),
			Error:  err.Error(),
		}, nil
	}

	return &pb.AdminSignUpResponse{
		Status: int64(status),
		Error:  "",
	}, nil

}

func (a *AdminServiceServer) AdminLogin(ctx context.Context, adminLoginDetails *pb.AdminLoginInRequest) (*pb.AdminLoginResponse, error) {

	adminLogin := models.AdminLogin{
		Email:    adminLoginDetails.Email,
		Password: adminLoginDetails.Password,
	}

	token, err := a.adminUseCase.AdminLogin(adminLogin)
	if err != nil {
		return &pb.AdminLoginResponse{
			Error: err.Error(),
		}, nil
	}

	return &pb.AdminLoginResponse{
		Token: token,
	}, nil

}

func (a *AdminServiceServer) ValidateAdmin(ctx context.Context, token *pb.AdminValidateRequest) (*pb.AdminValidateResponse, error) {

	signedToken := token.Token
	err := a.adminUseCase.ValidateAdmin(signedToken)
	if err != nil {
		return &pb.AdminValidateResponse{
			Status: http.StatusUnauthorized,
		}, err
	}

	return &pb.AdminValidateResponse{
		Status: http.StatusOK,
	}, nil
}
