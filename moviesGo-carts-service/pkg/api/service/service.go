package service

import (
	"context"

	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/cart/pb"
	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/domain"
	interfaces "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/usecase/interface"
)

type CartServiceServer struct {
	cartUseCase interfaces.CartUseCase
	pb.UnimplementedCartServiceServer
}

func NewCartServiceServer(useCase interfaces.CartUseCase) pb.CartServiceServer {

	return &CartServiceServer{
		cartUseCase: useCase,
	}

}


func (c *CartServiceServer) AddToCart(ctx context.Context, cartRequest *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {


	status,err :=  c.cartUseCase.AddToCart(int(cartRequest.Productid),int(cartRequest.Userid))

	if err != nil {
		return &pb.AddToCartResponse{
			Status: int64(status),
		},err
	}

	return &pb.AddToCartResponse{
		Status: int64(status),
	},nil

}

func (c *CartServiceServer) DisplayCart(ctx context.Context, userID *pb.DisplayCartRequest) (*pb.DisplayCartResponse, error) {
	
	cartResponse,err := c.cartUseCase.DisplayCArt(int(userID.Userid))
	if err != nil {
		return &pb.DisplayCartResponse{
		},err
	}

	var cartProducts []*pb.Cart
	for _,val := range cartResponse.Cart {
			cartProducts = append(cartProducts, &pb.Cart{
				Prdoductid: int64(val.ProductID),
				Moviename: val.MovieName,
				Quantity: int64(val.Quantity),
				Totalprice: float32(val.TotalPrice),
			})
	}

	return &pb.DisplayCartResponse{
		Totalprice: float32(cartResponse.TotalPrice),
		Cartproducts: cartProducts,
	},nil
	

}

func (c *CartServiceServer) OrderFromCart(ctx context.Context, orderRequest *pb.OrderRequest) (*pb.OrderResponse, error) {

	orderItems := []domain.OrderItems{}
	for _,product := range orderRequest.Cartproducts {
		orderItems = append(orderItems, domain.OrderItems{
			ProductID: int(product.Prdoductid),
			Quantity: int(product.Quantity),
			Price: float64(product.Totalprice),
		})
	}

	orderWrapper := domain.Order{
		UserID: int(orderRequest.Userid),
		AddressID: int(orderRequest.Addressid),
		Items: orderItems,
	}

	orderId, err := c.cartUseCase.OrderItemsFromCarts(orderWrapper)
	if err != nil {
		return &pb.OrderResponse{

		},err
	}

	return &pb.OrderResponse{
		Orderid: int64(orderId),
	},nil

}


	
