package handler

import (
	// 	"net/http"
	// 	"strconv"

	// "github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	interfaces "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/client/interface"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productClient interfaces.ProductClient
}

func NewProductHandler(productClientWrapper interfaces.ProductClient) *ProductHandler {
	return &ProductHandler{
		productClient: productClientWrapper,
	}
}

// // @Summary Get Products Details to users
// // @Description Retrieve all product Details with pagination to users
// // @Tags User Product
// // @Accept json
// // @Produce json
// // @Param page path string true "Page number"
// // @Param count query string true "Page Count"
// // @Success 200 {object} response.Response{}
// // @Failure 500 {object} response.Response{}
// // @Router /products/page/{page} [get]
// func (pr *ProductHandler) ShowAllProducts(c *gin.Context) {

// 	pageStr := c.Param("page")
// 	page, err := strconv.Atoi(pageStr)

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	count, err := strconv.Atoi(c.Query("count"))

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	products, err := pr.productUseCase.ShowAllProducts(page, count)

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not retrieve products", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products", products, nil)
// 	c.JSON(http.StatusOK, successRes)

// }

// @Summary Get Product Details To Admin
// @Description Retrieve product Details with pagination to Admin side
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param page path string true "Page number"
// @Param count query string true "Products Count Per Page"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products [get]
func (pr *ProductHandler) SeeAllProductToAdmin(c *gin.Context) {

	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)

	count, _ := strconv.Atoi(c.Query("count"))

	products, err := pr.productClient.ShowAllProducts(page,count)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products to admin side", products, nil)
	c.JSON(http.StatusOK, successRes)

}

// // @Summary Get Individual Product Details
// // @Description Get Individual Detailed product details to user side
// // @Tags User Product
// // @Accept json
// // @Produce json
// // @Param id path string true "sku"
// // @Success 200 {object} response.Response{}
// // @Failure 500 {object} response.Response{}
// // @Router /products/{id} [get]
// func (pr *ProductHandler) ShowIndividualProducts(c *gin.Context) {

// 	sku := c.Param("id")
// 	product, err := pr.productUseCase.ShowIndividualProducts(sku)

// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "path variables in wrong format", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Product details retrieved successfully", product, nil)
// 	c.JSON(http.StatusOK, successRes)

// }

// @Summary Add Products
// @Description Add product from admin side
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param product body models.ProductsReceiver true "Product details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products/add-product/ [post]
func (pr *ProductHandler) AddProduct(c *gin.Context) {

	var product models.ProductsReceiver

	if err := c.ShouldBindJSON(&product); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	productResponse, err := pr.productClient.AddProduct(product)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not add the product", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added products", productResponse, nil)
	c.JSON(http.StatusOK, successRes)

}
