package client

import (
	"context"
	"fmt"

	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/cart/pb"
	interfaces "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/client/interface"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/config"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"
	"google.golang.org/grpc"
)

type cartClient struct {
	cartClient pb.CartServiceClient
}

func NewCartClient(cfg config.Config) interfaces.CartClient {
	fmt.Println("cartsvc url = ",cfg.CartSvcUrl)
	grpcConnectoin, err := grpc.Dial(cfg.CartSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewCartServiceClient(grpcConnectoin)

	return &cartClient{
		cartClient: grpcClient,
	}

}

func (c *cartClient) AddToCart(productID int, userID int) (int, error) {

	res, err := c.cartClient.AddToCart(context.Background(), &pb.AddToCartRequest{
		Productid: int64(productID),
		Userid:    int64(userID),
	})
	fmt.Println("did error happen ath this point")

	if err != nil {
		return 400, err
	}

	return int(res.Status), nil

}

func (c *cartClient) DisplayCart(userID int) (models.CartResponse, error) {

	res, err := c.cartClient.DisplayCart(context.Background(),&pb.DisplayCartRequest{
		Userid: int64(userID),
	})

	if err != nil {
		return models.CartResponse{},err
	}

	var cartRes models.CartResponse
	cartRes.TotalPrice = float64(res.Totalprice)
	
	for _,carts := range res.Cartproducts {
		cartRes.Cart = append(cartRes.Cart, models.Cart{
			ProductID: uint(carts.Prdoductid),
			MovieName: carts.Moviename,
			Quantity: float64(carts.Quantity),
			TotalPrice: float64(carts.Totalprice),
		})
	}

	fmt.Println("DISPLAY CARTS: ",cartRes)

	return cartRes,nil

}
