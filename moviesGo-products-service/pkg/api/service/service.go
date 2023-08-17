package service

import (
	"context"
	"fmt"
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

func (p *ProductServiceServer) DoesProductExist(ctx context.Context, productID *pb.ProductExistRequest) (*pb.ProductExistResponse, error) {

	productExist, err := p.productUseCase.ProductExistInCarts(int(productID.Productid))
	if err != nil {
		return &pb.ProductExistResponse{
			Prouctexist: false,
		}, err
	}

	return &pb.ProductExistResponse{
		Prouctexist: productExist,
	}, nil
}

func (p *ProductServiceServer) GetProductPriceFromID(ctx context.Context, productID *pb.ProductPriceFromIDRequest) (*pb.ProductPriceFromIDResponse, error) {

	productPrice, err := p.productUseCase.GetProductPriceFromID(int(productID.Productid))
	if err != nil {
		return &pb.ProductPriceFromIDResponse{
			Price: 0.0,
		}, err
	}

	return &pb.ProductPriceFromIDResponse{
		Price: float32(productPrice),
	}, nil

}

func (p *ProductServiceServer) GetProductNameFromID(context.Context, *pb.ProductNameFromIDRequest) (*pb.ProductNameFromIDResponse, error) {

	return &pb.ProductNameFromIDResponse{
		Productname: "okokok",
	}, nil

}

func (p *ProductServiceServer) GetCartProductsNameFromID(ctx context.Context, productID *pb.ProductIDSRequest) (*pb.ProductNamesResponse, error) {

	var productIDS []int

	for _, p := range productID.Allproductids {
		fmt.Println("product ID associated with this = : ", p.Productid)
		productIDS = append(productIDS, int(p.Productid))
	}

	fmt.Println("array of product ID's : ", productIDS)

	productMap, err := p.productUseCase.GetProductsNameFromID(productIDS)
	if err != nil {
		return &pb.ProductNamesResponse{}, err
	}

	var productNamesRes []*pb.ProductNameWithID

	for _, p := range productMap {
		productNamesRes = append(productNamesRes, &pb.ProductNameWithID{
			Productid:   int64(p.ID),
			Productname: p.MovieName,
		})
	}

	return &pb.ProductNamesResponse{
		Productswithid: productNamesRes,
	}, nil

}
