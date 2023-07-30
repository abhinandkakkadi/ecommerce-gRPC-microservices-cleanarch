package service

import (
	"context"
	"strconv"

	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/product/pb"
	interfaces "github.com/abhinandkakkadi/moviesgo-products-service/pkg/usecase/interface"
	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/utils/models"
)

type ProductServiceServer struct {
	productUseCase interfaces.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func NewProductServiceServer(useCase interfaces.ProductUseCase) pb.ProductServiceServer {

	return &ProductServiceServer{
		productUseCase: useCase,
	}

}

func (p *ProductServiceServer) ShowAllProducts(ctx context.Context, productRequest *pb.ShowProductRequest) (*pb.ShowAllProductResponse, error) {

	count, page := &productRequest.Count, &productRequest.Page
	productsResponse, err := p.productUseCase.ShowAllProducts(int(*page), int(*count))

	var products []*pb.ShowProductResponse
	for _, p := range productsResponse {
		products = append(products, &pb.ShowProductResponse{
			Productid: int64(p.ID),
			Moviename: p.MovieName,
			Sku:       p.Sku,
			Genre:     p.Genre,
			Price:     strconv.FormatFloat(p.Price, 'f', -1, 64),
			Language:  p.Language,
			Quantity:  int64(p.Quantity),
		})
	}

	if err != nil {
		return &pb.ShowAllProductResponse{
			Productsbrief: nil,
		}, err
	}

	return &pb.ShowAllProductResponse{
		Productsbrief: products,
	}, nil

}

func (p *ProductServiceServer) AddProduct(ctx context.Context, product *pb.AddProductRequest) (*pb.AddProductResponse, error) {

	productsReceiver := models.ProductsReceiver{
		MovieName:           product.Moviename,
		GenreID:             uint(product.Genreid),
		ReleaseYear:         product.Releaseyear,
		Format:              product.Format,
		Director:            product.Director,
		ProductsDescription: product.Productdescription,
		Runtime:             float64(product.Runtime),
		Language:            product.Language,
		StudioID:            uint(product.Studioid),
		Quantity:            int(product.Quantity),
		Price:               float64(product.Price),
	}

	_, err := p.productUseCase.AddProduct(productsReceiver)
	if err != nil {
		return &pb.AddProductResponse{
			Status: 400,
		}, err
	}

	return &pb.AddProductResponse{
		Status: 201,
	}, nil

}

func (p *ProductServiceServer) GetGenreDetails(context.Context, *pb.GenreRequest) (*pb.AllGenreResponse, error) {

	return &pb.AllGenreResponse{}, nil
}

func (p *ProductServiceServer) GetStudioDetails(context.Context, *pb.StudioRequest) (*pb.AllStudioResponse, error) {
	return &pb.AllStudioResponse{}, nil
}
