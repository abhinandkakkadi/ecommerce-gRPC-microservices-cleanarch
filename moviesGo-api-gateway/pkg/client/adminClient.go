package client

import (
	"context"
	"fmt"

	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/admin/pb"
	interfaces "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/client/interface"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/config"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"
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



func (a *adminClient) LoginHandler(adminDetails models.AdminLogin) (string,error) {

	res, err := a.adminClient.AdminLogin(context.Background(),&pb.LoginInRequest{
		Email: adminDetails.Email,
		Password: adminDetails.Password,
	})

	if err != nil {
		return "",err
	}

	return res.Token,nil

}


func (a *adminClient) CreateAdmin(admin models.AdminSignUp) (int,error) {

	res, err := a.adminClient.AdminSignUp(context.Background(),&pb.SingUpRequest{
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