package handler

// import (
// 	"net/http"
// 	"strconv"

// 	services "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/usecase/interface"
// 	models "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"
// 	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/response"
// 	"github.com/gin-gonic/gin"
// )

// type OrderHandler struct {
// 	orderUseCase services.OrderUseCase
// }

// func NewOrderHandler(useCase services.OrderUseCase) *OrderHandler {
// 	return &OrderHandler{
// 		orderUseCase: useCase,
// 	}
// }

// // @Summary Order Items from cart
// // @Description Order all products which is currently present inside  the cart
// // @Tags User Order
// // @Accept json
// // @Produce json
// // @Security Bearer
// // @Param orderBody body models.OrderFromCart true "Order details"
// // @Success 200 {object} response.Response{}
// // @Failure 500 {object} response.Response{}
// // @Router /order [post]
// func (o *OrderHandler) OrderItemsFromCart(c *gin.Context) {

// 	id, _ := c.Get("user_id")
// 	userID := id.(int)

// 	var orderFromCart models.OrderFromCart
// 	if err := c.ShouldBindJSON(&orderFromCart); err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "bad request", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	orderSuccessResponse, err := o.orderUseCase.OrderItemsFromCart(orderFromCart, userID)
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Successfully created the order", orderSuccessResponse, nil)
// 	c.JSON(http.StatusOK, successRes)

// }

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
// func (o *OrderHandler) GetOrderDetails(c *gin.Context) {

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

// 	fullOrderDetails, err := o.orderUseCase.GetOrderDetails(userID, page, pageSize)

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Full Order Details", fullOrderDetails, nil)
// 	c.JSON(http.StatusOK, successRes)

// }



// // @Summary Get All order details for admin
// // @Description Get all order details to the admin side
// // @Tags Admin Order Management
// // @Accept json
// // @Produce json
// // @Security Bearer
// // @Param page path string true "Page number"
// // @Success 200 {object} response.Response{}
// // @Failure 500 {object} response.Response{}
// // @Router /admin/orders/{id} [get]
// func (o *OrderHandler) GetAllOrderDetailsForAdmin(c *gin.Context) {

// 	pageStr := c.Param("page")
// 	page, err := strconv.Atoi(pageStr)

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	allOrderDetails, err := o.orderUseCase.GetAllOrderDetailsForAdmin(page)

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not retrieve order details", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Order Details Retrieved successfully", allOrderDetails, nil)
// 	c.JSON(http.StatusOK, successRes)

// }


