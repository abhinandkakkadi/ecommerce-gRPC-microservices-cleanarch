// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/api"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/api/handler"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/client"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/config"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {

	cartClient := client.NewCartClient(cfg)
	cartHandler := handler.NewCartHandler(cartClient)

	userClient := client.NewUserClient(cfg, cartClient)
	userHandler := handler.NewUserHandler(userClient)

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	productClient := client.NewProductClient(cfg)
	productHandler := handler.NewProductHandler(productClient)

	// orderHandler := handler.NewOrderHandler(orderUseCase)

	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, userClient, adminClient, productHandler, cartHandler)

	return serverHTTP, nil
}
