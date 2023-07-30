//go:build wireinject
// +build wireinject

package di

import (
	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/api"
	http "github.com/abhinandkakkadi/moviesgo-products-service/pkg/api"
	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/api/service"
	config "github.com/abhinandkakkadi/moviesgo-products-service/pkg/config"
	db "github.com/abhinandkakkadi/moviesgo-products-service/pkg/db"
	repository "github.com/abhinandkakkadi/moviesgo-products-service/pkg/repository"
	usecase "github.com/abhinandkakkadi/moviesgo-products-service/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewAdminRepository,usecase.NewAdminUseCase,service.NewAdminServiceServer,api.NewGRPCServer)
	return &http.ServerHTTP{}, nil
}


