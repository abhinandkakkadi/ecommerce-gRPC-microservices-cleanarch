package service

import (
	"context"
	"fmt"

	interfaces "github.com/abhinandkakkadi/moviesgo-user-service/pkg/usecase/interface"
	"github.com/abhinandkakkadi/moviesgo-user-service/pkg/user/pb"
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


func (u *UserServiceServer) SampleRequest(ctx context.Context,sampleRequest *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	fmt.Println("health check passed")
	return &pb.RegisterResponse{
		Response: "Abhinand",
	},nil

}

func (u *UserServiceServer) UserSignUp(ctx context.Context,userDetails *pb.SingUpRequest) (*pb.SignUpResponse, error) {

	fmt.Println("Reached SignUp handler")
	
}
