package handler

// import (
// 	"net/http"

// 	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"
// 	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/response"
// 	"github.com/gin-gonic/gin"
// )

// type AdminHandler struct {
// 	adminUseCase services.AdminUseCase
// }

// func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
// 	return &AdminHandler{
// 		adminUseCase: usecase,
// 	}
// }

// // @Summary Admin Login
// // @Description Login handler for admin
// // @Tags Admin Authentication
// // @Accept json
// // @Produce json
// // @Param  admin body models.AdminLogin true "Admin login details"
// // @Success 200 {object} response.Response{}
// // @Failure 500 {object} response.Response{}
// // @Router /admin/adminlogin [post]
// func (cr *AdminHandler) LoginHandler(c *gin.Context) { // login handler for the admin

// 	var adminDetails models.AdminLogin
// 	if err := c.ShouldBindJSON(&adminDetails); err != nil {
// 		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}

// 	admin, err := cr.adminUseCase.LoginHandler(adminDetails)
// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
// 	c.JSON(http.StatusOK, successRes)

// }

// // @Summary Admin Signup
// // @Description Signup handler for admin
// // @Tags Admin Authentication
// // @Accept json
// // @Produce json
// // @Security Bearer
// // @Param  admin body models.AdminSignUp true "Admin login details"
// // @Success 200 {object} response.Response{}
// // @Failure 500 {object} response.Response{}
// // @Router /admin/createadmin [post]
// func (cr *AdminHandler) CreateAdmin(c *gin.Context) {

// 	var admin models.AdminSignUp
// 	if err := c.ShouldBindJSON(&admin); err != nil {
// 		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errRes)
// 	}

// 	adminDetails, err := cr.adminUseCase.CreateAdmin(admin)
// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
// 		c.JSON(http.StatusInternalServerError, errRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusCreated, "Successfully signed up the user", adminDetails, nil)
// 	c.JSON(http.StatusCreated, successRes)

// }
