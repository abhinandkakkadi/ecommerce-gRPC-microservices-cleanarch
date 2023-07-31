package client

import (
	"context"
	"fmt"
	"strconv"

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
	fmt.Println("product client: ",cfg.ProductSvcUrl)
	grpcConnectoin, err := grpc.Dial(cfg.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewProductServiceClient(grpcConnectoin)

	return &productClient{
		client: grpcClient,
	}

}

func (p *productClient) ShowAllProducts(page int, count int) ([]models.ProductsBrief, error) {

	productDetails, err := p.client.ShowAllProducts(context.Background(), &pb.ShowProductRequest{
		Page:  int64(page),
		Count: int64(count),
	})

	if err != nil {
		return []models.ProductsBrief{}, err
	}

	var products []models.ProductsBrief
	for _, p := range productDetails.Productsbrief {
		price, _ := strconv.ParseFloat(p.Price, 64)
		products = append(products, models.ProductsBrief{
			ID:        int(p.Productid),
			MovieName: p.Moviename,
			Sku:       p.Sku,
			Genre:     p.Genre,
			Language:  p.Language,
			Price:     price,
			Quantity:  int(p.Quantity),
		})
	}

	return products, nil

}

func (p *productClient) AddProduct(productDetails models.ProductsReceiver) (int, error) {

	status, err := p.client.AddProduct(context.Background(), &pb.AddProductRequest{
		Moviename:          productDetails.MovieName,
		Genreid:            int64(productDetails.GenreID),
		Releaseyear:        productDetails.ReleaseYear,
		Format:             productDetails.Format,
		Director:           productDetails.Director,
		Productdescription: productDetails.ProductsDescription,
		Runtime:            int64(productDetails.Runtime),
		Language:           productDetails.Language,
		Studioid:           int64(productDetails.StudioID),
		Quantity:           int64(productDetails.Quantity),
		Price:              int64(productDetails.Price),
	})

	if err != nil {
		return 0, nil
	}

	return int(status.Status), nil

}

// func (p *productClient) GetGenreDetails() {

// }
// func (p *productClient) GetStudioDetails() {

// }
