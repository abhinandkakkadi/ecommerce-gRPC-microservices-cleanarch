package service

import (
	"context"

	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/cart/pb"
	interfaces "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/usecase/interface"
)

type CartServiceServer struct {
	cartUseCase interfaces.CartUseCase
	pb.UnimplementedCartServiceServer
}

func NewCartServiceServer(useCase interfaces.CartUseCase) pb.CartServiceServer {

	return &CartServiceServer{
		cartUseCase: useCase,
	}

}


func (c *CartServiceServer) AddToCart(ctx context.Context, cartRequest *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {


	c.cartUseCase.AddToCart(int(cartRequest.Productid),int(cartRequest.Userid))
	return &pb.AddToCartResponse{
		Status: 400,
	},nil

}

func DisplayCart(context.Context, *pb.DisplayCartRequest) (*pb.DisplayCartResponse, error) {
	return &pb.DisplayCartResponse{},nil
}


	
