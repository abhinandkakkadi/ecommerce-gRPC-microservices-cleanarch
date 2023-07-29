//go:build wireinject
// +build wireinject

package di

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	http "github.com/abhinandkakkadi/moviesgo-user-service/pkg/api"
	config "github.com/abhinandkakkadi/moviesgo-user-service/pkg/config"
	db "github.com/abhinandkakkadi/moviesgo-user-service/pkg/db"
	repository "github.com/abhinandkakkadi/moviesgo-user-service/pkg/repository"
	usecase "github.com/abhinandkakkadi/moviesgo-user-service/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP, repository.NewProductRepository, usecase.NewProductUseCase, handler.NewProductHandler, handler.NewOtpHandler, usecase.NewOtpUseCase, repository.NewOtpRepository, repository.NewAdminRepository, usecase.NewAdminUseCase, handler.NewAdminHandler, handler.NewCartHandler, usecase.NewCartUseCase, repository.NewCartRepository)
	return &http.ServerHTTP{}, nil
}
