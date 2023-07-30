package routes

import (
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/api/handler"
	interfaces "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/client/interface"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, adminHandler *handler.AdminHandler, adminClient interfaces.AdminClient, productHandler *handler.ProductHandler) {

	router.POST("/adminlogin", adminHandler.LoginHandler)
	router.POST("/create-admin",adminHandler.CreateAdmin)

	router.Use(adminClient.AdminAuthRequired)
	{

		product := router.Group("/products")
		{
			// product.GET("", productHandler.ShowAllProductsToAdmin)
			// product.GET("/:page", productHandler.ShowAllProductToAdmin)
			product.POST("/add-product", productHandler.AddProduct)

		}

	}

	// orders := router.Group("/orders")
	// 	{
	// 		orders.GET("", orderHandler.GetAllOrderDetailsForAdmin)
	// 		orders.GET("/:page", orderHandler.GetAllOrderDetailsForAdmin)

	// 	}

}
