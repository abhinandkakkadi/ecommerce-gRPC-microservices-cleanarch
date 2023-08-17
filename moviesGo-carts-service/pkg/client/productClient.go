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

func (p *productClient) DoesProductExist(productID int) (bool, error) {

	res, err := p.client.DoesProductExist(context.Background(), &pb.ProductExistRequest{
		Productid: int64(productID),
	})

	if err != nil {
		return false, err
	}

	return res.Prouctexist, nil

}

func (p *productClient) GetProductPriceFromID(productID int) (float64, error) {

	res, err := p.client.GetProductPriceFromID(context.Background(), &pb.ProductPriceFromIDRequest{
		Productid: int64(productID),
	})

	if err != nil {
		return 0.0, err
	}

	return float64(res.Price), nil

}

func (p *productClient) GetProductNamesFromID(productId []int) (map[int]string, error) {

	var productMap []*pb.ProductIDS

	for _, p := range productId {
		productMap = append(productMap, &pb.ProductIDS{
			Productid: int64(p),
		})
	}

	res, err := p.client.GetCartProductsNameFromID(context.Background(), &pb.ProductIDSRequest{
		Allproductids: productMap,
	})

	if err != nil {
		return map[int]string{}, nil
	}

	products := make(map[int]string)
	for _, p := range res.Productswithid {
		fmt.Println("PRODUCT NAME: ", p.Productname)
		products[int(p.Productid)] = p.Productname
	}

	return products, nil

}
