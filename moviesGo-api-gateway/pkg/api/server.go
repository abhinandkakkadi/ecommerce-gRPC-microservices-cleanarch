package http

import (
	"log"

	handler "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/api/handler"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/api/routes"
	interfaces "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/client/interface"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,adminHandler *handler.AdminHandler,userClient interfaces.UserClient,adminClient interfaces.AdminClient,productHandler *handler.ProductHandler) *ServerHTTP {
	router := gin.New()

	router.Use(gin.Logger())

	routes.UserRoutes(router.Group("/"), userHandler,userClient)
	routes.AdminRoutes(router.Group("/admin"), adminHandler,adminClient,productHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	log.Printf("starting server on :3000")
	err := sh.engine.Run(":3000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
