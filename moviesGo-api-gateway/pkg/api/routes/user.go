package routes

import (
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/api/handler"
	interfaces "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/client/interface"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, userClient interfaces.UserClient, productHandler *handler.ProductHandler, cartHandler *handler.CartHandler) {

	// // USER SIDE
	router.POST("/signup", userHandler.UserSignUp)
	router.POST("/login", userHandler.LoginHandler)

	router.Use(userClient.UserAuthRequired)
	router.GET("/sample", userHandler.SampleRequest)
	product := router.Group("/products")

	{
		product.GET("", productHandler.ShowAllProductToAdmin)
		product.GET("/page/:page", productHandler.ShowAllProductToAdmin)
	}

	cart := router.Group("/cart")
	{
		cart.PATCH("/addtocart/:id", cartHandler.AddToCart)
		cart.GET("", cartHandler.DisplayCart)
	}

	router.GET("/checkout", userHandler.CheckOut)
	router.POST("/order", cartHandler.OrderItemsFromCart)

	// }

	users := router.Group("/users")
	{
		users.GET("/address", userHandler.GetAllAddress)
		users.POST("/address", userHandler.AddAddress)
		// orders := users.Group("orders")
		// {
		// 	orders.GET("", orderHandler.GetOrderDetails)
		// 	orders.GET("/:page", orderHandler.GetOrderDetails)
		// }

	}
}
