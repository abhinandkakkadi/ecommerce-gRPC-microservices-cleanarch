package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/admin/pb"
	interfaces "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/client/interface"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/config"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)


type adminClient struct {
	adminClient pb.AdminAuthServiceClient
}

func NewAdminClient(cfg config.Config) interfaces.AdminClient {

	grpcConnectoin, err := grpc.Dial(cfg.AdminAuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewAdminAuthServiceClient(grpcConnectoin)

	return &adminClient{
		adminClient: grpcClient,
	}

}

func (a *adminClient) AdminAuthRequired(c *gin.Context) {

	token := c.GetHeader("authorization")

	res,err := a.adminClient.ValidateAdmin(context.Background(),&pb.AdminValidateRequest{
		Token: token,
	})

	if err != nil {
		c.AbortWithStatus(int(res.Status))
		errRes := response.ClientResponse(http.StatusUnauthorized, "client unauthorized to access this route", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	

	c.Next()

}



func (a *adminClient) LoginHandler(adminDetails models.AdminLogin) (string,error) {

	res, _ := a.adminClient.AdminLogin(context.Background(),&pb.AdminLoginInRequest{
		Email: adminDetails.Email,
		Password: adminDetails.Password,
	})

	if res.Error != "" {
		return "",errors.New(res.Error)
	}

	return res.Token,nil

}


func (a *adminClient) CreateAdmin(admin models.AdminSignUp) (int,error) {

	res, err := a.adminClient.AdminSignUp(context.Background(),&pb.AdminSingUpRequest{
		Name: admin.Name,
		Email: admin.Email,
		Password: admin.Password,
		Confirmpassword: admin.ConfirmPassword,
	})

	if err != nil {
		return 400,err
	}

	return int(res.Status),nil

}