package service

import (
	"context"

	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/product/pb"
	interfaces "github.com/abhinandkakkadi/moviesgo-products-service/pkg/usecase/interface"
)

type ProductServiceServer struct {
	adminUseCase interfaces.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func NewProductServiceServer(useCase interfaces.ProductUseCase) pb.ProductServiceServer {

	return &ProductServiceServer{
		adminUseCase: useCase,
	}

}

func (p *ProductServiceServer) ShowAllProducts(context.Context, *pb.ShowProductRequest) (*pb.ShowAllProductResponse, error) {

}
	

func (p *ProductServiceServer) AddProduct(context.Context, *pb.AddProductRequest) (*pb.AddProductResponse, error) {

}


func (p *ProductServiceServer) GetGenreDetails(context.Context, *pb.GenreRequest) (*pb.AllGenreResponse, error) {

}


func (p *ProductServiceServer) GetStudioDetails(context.Context, *pb.StudioRequest) (*pb.AllStudioResponse, error) {

}

	
	

