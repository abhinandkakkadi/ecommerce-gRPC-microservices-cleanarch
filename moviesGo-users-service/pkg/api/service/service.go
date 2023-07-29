package service

import (
	"context"

	"github.com/abhinandkakkadi/moviesgo-user-service/pkg/pb"
)

type UserServiceServer struct {
	// userUseCase interfaces.UserUseCase
	pb.UnimplementedAuthServiceServer
}


func NewUserServiceServer() pb.AuthServiceServer {

	return &UserServiceServer{
		// userUseCase: useCase,
	}

}


func (u *UserServiceServer) SampleRequest(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	return &pb.RegisterResponse{
		Response: "Abhinand",
	},nil

}
