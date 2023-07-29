package routes

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/handler"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler, productHandler *handler.ProductHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler) {

	// USER SIDE
	router.POST("/signup", userHandler.UserSignUp)
	router.POST("/login", userHandler.LoginHandler)


	product := router.Group("/products")
	{
		product.GET("", productHandler.ShowAllProducts)
		product.GET("/page/:page", productHandler.ShowAllProducts)
		product.GET("/:id", productHandler.ShowIndividualProducts)

	}

	router.Use(middleware.AuthMiddleware())

	{
		cart := router.Group("/cart")
		{
			cart.POST("/addtocart/:id", cartHandler.AddToCart)
			cart.DELETE("/removefromcart/:id", cartHandler.RemoveFromCart)
			cart.GET("", cartHandler.DisplayCart)
			cart.DELETE("", cartHandler.EmptyCart)

		}

		router.GET("/checkout", userHandler.CheckOut)
		router.POST("/order", orderHandler.OrderItemsFromCart)

	}

	users := router.Group("/users")
		{
			users.GET("/address", userHandler.GetAllAddress)
			users.POST("/address", userHandler.AddAddress)
			orders := users.Group("orders")
			{
				orders.GET("", orderHandler.GetOrderDetails)
				orders.GET("/:page", orderHandler.GetOrderDetails)
			}

		}
}
