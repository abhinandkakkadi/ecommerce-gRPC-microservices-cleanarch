package client

import (
	"context"
	"fmt"

	interfaces "github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/client/interface"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/config"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/product/pb"
	"github.com/abhinandkakkadi/moviesgo-api-gateway/pkg/utils/models"
	"google.golang.org/grpc"
)




type productClient struct {
	client pb.ProductServiceClient
}

func NewProductClient(cfg config.Config) interfaces.ProductClient {

	grpcConnectoin, err := grpc.Dial(cfg.AdminAuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewProductServiceClient(grpcConnectoin)

	return &productClient{
		client: grpcClient,
	}

}

func (p *productClient) ShowAllProducts(page int, count int) (models.ProductsBrief,error) {

	productDetails,err :=  p.client.ShowAllProducts(context.Background(),&pb.ShowProductRequest{
		Page: int64(page),
		Count: int64(count),
	})

	if err != nil {
		return models.ProductsBrief{},err
	}

	_ = productDetails

	return models.ProductsBrief{},nil
	
}

func (p *productClient) AddProduct(productDetails models.ProductsReceiver) (int,error) {

	status, err := p.client.AddProduct(context.Background(),&pb.AddProductRequest{
		Moviename: productDetails.MovieName,
		Genreid: int64(productDetails.GenreID),
		Releaseyear: productDetails.ReleaseYear,
		Format: productDetails.Format,
		Director: productDetails.Director,
		Productdescription: productDetails.ProductsDescription,
		Runtime: int64(productDetails.Runtime),
		Language: productDetails.Language,
		Studioid: int64(productDetails.StudioID),
		Quantity: int64(productDetails.Quantity),
		Price: int64(productDetails.Price),
	})

	_, _  = status, err
	return 0,nil
	
}
// func (p *productClient) GetGenreDetails() {

// } 
// func (p *productClient) GetStudioDetails() {

// } 