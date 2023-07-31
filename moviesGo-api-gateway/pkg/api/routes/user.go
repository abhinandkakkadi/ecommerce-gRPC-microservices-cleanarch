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
		// product.GET("/:id", productHandler.ShowIndividualProducts)

	}

	cart := router.Group("/cart")
	{
		cart.GET("/addtocart/:id", cartHandler.AddToCart)
		// cart.DELETE("/removefromcart/:id", cartHandler.RemoveFromCart)
		cart.GET("", cartHandler.DisplayCart)
		// cart.DELETE("", cartHandler.EmptyCart)

	}

	// router.Use(middleware.AuthMiddleware())

	// {
	// cart := router.Group("/cart")
	// {
	// 	cart.POST("/addtocart/:id", cartHandler.AddToCart)
	// 	cart.DELETE("/removefromcart/:id", cartHandler.RemoveFromCart)
	// 	cart.GET("", cartHandler.DisplayCart)
	// 	cart.DELETE("", cartHandler.EmptyCart)

	// }

	// 	router.GET("/checkout", userHandler.CheckOut)
	// 	router.POST("/order", orderHandler.OrderItemsFromCart)

	// }

	// users := router.Group("/users")
	// 	{
	// 		users.GET("/address", userHandler.GetAllAddress)
	// 		users.POST("/address", userHandler.AddAddress)
	// 		orders := users.Group("orders")
	// 		{
	// 			orders.GET("", orderHandler.GetOrderDetails)
	// 			orders.GET("/:page", orderHandler.GetOrderDetails)
	// 		}

	// 	}
}
