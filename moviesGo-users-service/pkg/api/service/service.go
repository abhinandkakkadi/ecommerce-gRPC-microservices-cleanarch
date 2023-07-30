package service

import (
	"context"
	"fmt"
	"net/http"

	interfaces "github.com/abhinandkakkadi/moviesgo-user-service/pkg/usecase/interface"
	"github.com/abhinandkakkadi/moviesgo-user-service/pkg/user/pb"
	"github.com/abhinandkakkadi/moviesgo-user-service/pkg/utils/models"
)

type UserServiceServer struct {
	userUseCase interfaces.UserUseCase
	pb.UnimplementedAuthServiceServer
}

func NewUserServiceServer(useCase interfaces.UserUseCase) pb.AuthServiceServer {

	return &UserServiceServer{
		userUseCase: useCase,
	}

}

func (u *UserServiceServer) SampleRequest(ctx context.Context, sampleRequest *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	fmt.Println("health check passed")
	return &pb.RegisterResponse{
		Response: "Abhinand",
	}, nil

}

func (u *UserServiceServer) UserSignUp(ctx context.Context, userDetails *pb.SingUpRequest) (*pb.SignUpResponse, error) {

	fmt.Println("Reached SignUp handler")
	userSignUpDetails := models.UserDetails{
		Name:            userDetails.Name,
		Email:           userDetails.Email,
		Phone:           userDetails.Phone,
		Password:        userDetails.Password,
		ConfirmPassword: userDetails.Confirmpassword,
	}

	status, err := u.userUseCase.UserSignUp(userSignUpDetails)
	if err != nil {
		return &pb.SignUpResponse{
			Status: int64(status),
			Error:  err.Error(),
		}, nil

	}

	return &pb.SignUpResponse{
		Status: int64(status),
	}, nil

}

func (u *UserServiceServer) UserLogin(ctx context.Context, user *pb.LoginInRequest) (*pb.LoginResponse, error) {

	userLoginDetails := models.UserLogin{
		Email:    user.Email,
		Password: user.Password,
	}

	token, err := u.userUseCase.UserLogin(userLoginDetails)
	if err != nil {
		return &pb.LoginResponse{
			Token: "",
			Error: err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Token: token,
		Error: "",
	}, nil

}

func (u *UserServiceServer) ValidateUser(ctx context.Context, token *pb.ValidateRequest) (*pb.ValidateResponse, error) {

	signedToken := token.Token
	userID, err := u.userUseCase.ValidateUser(signedToken)
	if err != nil {
		return &pb.ValidateResponse{
			UserID: 0,
			Status: http.StatusUnauthorized,
		}, err
	}

	return &pb.ValidateResponse{
		UserID: int64(userID),
		Status: http.StatusOK,
	}, nil

}
