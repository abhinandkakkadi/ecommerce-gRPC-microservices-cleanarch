package client

import (
	"context"
	"fmt"

	interfaces "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/client/interface"
	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/config"
	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/product/pb"
	"google.golang.org/grpc"
)

type productClient struct {
	client pb.ProductServiceClient
}

func NewProductClient(cfg config.Config) interfaces.ProductClient {

	grpcConnection, err := grpc.Dial(cfg.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewProductServiceClient(grpcConnection)

	return &productClient{
		client: grpcClient,
	}

}


func (p *productClient) DoesProductExist(productID int) (bool,error) {

	res, err := p.client.DoesProductExist(context.Background(),&pb.ProductExistRequest{
		Productid: int64(productID),
	})

	if err != nil {
		return false,err
	}

	return res.Result,nil

}

