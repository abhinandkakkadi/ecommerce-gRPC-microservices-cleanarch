package usecase

import (
	interfaces "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/repository/interface"
	services "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/usecase/interface"
)

type cartUseCase struct {
	cartRepository interfaces.CartRepository
}

func NewCartUseCase(repo interfaces.CartRepository) services.CartUseCase {

	return &cartUseCase{
		cartRepository: repo,
	}
}

