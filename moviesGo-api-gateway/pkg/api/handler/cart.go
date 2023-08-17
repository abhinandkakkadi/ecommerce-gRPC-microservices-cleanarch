package handler

import (
	// "net/http"
	// "strconv"
	"fmt"
	"net/http"
	"strconv"

	services "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/client/interface"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
	// services "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/usecase/interface"
	// "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/response"
	// "github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartClient services.CartClient
}

func NewCartHandler(client services.CartClient) *CartHandler {

	return &CartHandler{
		cartClient: client,
	}

}

// @Summary Add to Cart
// @Description Add product to the cart using product id
// @Tags User Cart
// @Accept json
// @Produce json
// @Param id path string true "product-id"
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart/addtocart/{id} [post]
func (cr *CartHandler) AddToCart(c *gin.Context) {

	id := c.Param("id")
	product_id, err := strconv.Atoi(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "product id not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	userID, _ := c.Get("userID")
	fmt.Println("UserID value : ", userID)
	cartResponse, err := cr.cartClient.AddToCart(product_id, int(userID.(int64)))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not add product to the cart", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "successfully added product to the cart", cartResponse, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary Remove product from cart
// @Description Remove specified product of quantity 1 from cart using product id
// @Tags User Cart
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart/removefromcart/{id} [delete]
// func (cr *CartHandler) RemoveFromCart(c *gin.Context) {

// 	id := c.Param("id")
// 	product_id, err := strconv.Atoi(id)

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "product not in right format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	userID, _ := c.Get("user_id")
// 	updatedCart, err := cr.cartUseCase.RemoveFromCart(product_id, userID.(int))

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not delete cart items", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Cart item deleted", updatedCart, nil)
// 	c.JSON(http.StatusOK, successRes)

// }

// @Summary Display Cart
// @Description Display all products of the cart along with price of the product and grand total
// @Tags User Cart
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart [get]
func (cr *CartHandler) DisplayCart(c *gin.Context) {

	userID, _ := c.Get("userID")
	cart, err := cr.cartClient.DisplayCart(int(userID.(int64)))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not display cart items", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Cart items displayed successfully", cart, nil)
	c.JSON(http.StatusOK, successRes)

}

// // @Summary Delete all Items Present inside the Cart
// // @Description Remove all product from cart
// // @Tags User Cart
// // @Accept json
// // @Produce json
// // @Security Bearer
// // @Success 200 {object} response.Response{}
// // @Failure 500 {object} response.Response{}
// // @Router /cart [delete]
// func (cr *CartHandler) EmptyCart(c *gin.Context) {

// 	userID, _ := c.Get("user_id")
// 	emptyCart, err := cr.cartUseCase.EmptyCart(userID.(int))

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not display cart items", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "cart items displayed successfully", emptyCart, nil)
// 	c.JSON(http.StatusOK, successRes)
// }

// @Summary Order Items from cart
// @Description Order all products which is currently present inside  the cart
// @Tags User Order
// @Accept json
// @Produce json
// @Security Bearer
// @Param orderBody body models.OrderFromCart true "Order details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /order [post]
func (o *CartHandler) OrderItemsFromCart(c *gin.Context) {

	id, _ := c.Get("userID")
	userID := id.(int64)

	var orderFromCart models.OrderFromCart
	if err := c.ShouldBindJSON(&orderFromCart); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orderSuccessResponse, err := o.cartClient.OrderItemsFromCart(orderFromCart, int(userID))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully created the order", orderSuccessResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

// // @Summary Get Order Details to user side
// // @Description Get all order details done by user to user side
// // @Tags User Order
// // @Accept json
// // @Produce json
// // @Security Bearer
// // @Param id path string true "page number"
// // @Param pageSize query string true "page size"
// // @Success 200 {object} response.Response{}
// // @Failure 500 {object} response.Response{}
// // @Router /users/orders/{id} [get]
// func (o *CartHandler) GetOrderDetails(c *gin.Context) {

// 	pageStr := c.Param("page")
// 	page, err := strconv.Atoi(pageStr)

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in correct format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	pageSize, err := strconv.Atoi(c.Query("count"))

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	id, _ := c.Get("user_id")
// 	userID := id.(int)

// 	fullOrderDetails, err := o.cartClient.GetOrderDetails(userID, page, pageSize)

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Full Order Details", fullOrderDetails, nil)
// 	c.JSON(http.StatusOK, successRes)

// }
