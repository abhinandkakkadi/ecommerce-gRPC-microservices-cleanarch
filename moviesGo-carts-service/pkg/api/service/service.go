package service

import (
	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/cart/pb"
	interfaces "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/usecase/interface"
)

type CartServiceServer struct {
	productUseCase interfaces.CartUseCase
	pb.UnimplementedCartServiceServer
}

func NewCartServiceServer(useCase interfaces.CartUseCase) pb.CartServiceServer {

	return &CartServiceServer{
		productUseCase: useCase,
	}

}

