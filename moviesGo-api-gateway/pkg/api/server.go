package http

import (
	"log"

	handler "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/api/handler"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/api/routes"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,adminHandler *handler.AdminHandler) *ServerHTTP {
	router := gin.New()

	router.Use(gin.Logger())

	routes.UserRoutes(router.Group("/"), userHandler)
	routes.AdminRoutes(router.Group("/admin"), adminHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	log.Printf("starting server on :3000")
	err := sh.engine.Run(":3000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
