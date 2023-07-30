package handler

import (
	// "net/http"
	// "strconv"
	"net/http"
	"strconv"

	services "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/client/interface"
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

	userID, _ := c.Get("user_id")
	cartResponse, err := cr.cartClient.AddToCart(product_id, userID.(int))

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

	userID, _ := c.Get("user_id")
	cart, err := cr.cartClient.DisplayCart(userID.(int))

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
