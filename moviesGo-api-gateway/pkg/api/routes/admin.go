package routes

import (
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, adminHandler *handler.AdminHandler) {

	router.POST("/adminlogin", adminHandler.LoginHandler)
	router.POST("/create-admin",adminHandler.CreateAdmin)

	// router.Use(middleware.AuthorizationMiddleware)
	// {

	// 	product := router.Group("/products")
	// 	{
	// 		product.GET("", productHandler.SeeAllProductToAdmin)
	// 		product.GET("/:page", productHandler.SeeAllProductToAdmin)
	// 		product.POST("/add-product", productHandler.AddProduct)

	// 	}

	// }

	// orders := router.Group("/orders")
	// 	{
	// 		orders.GET("", orderHandler.GetAllOrderDetailsForAdmin)
	// 		orders.GET("/:page", orderHandler.GetAllOrderDetailsForAdmin)

	// 	}

}
