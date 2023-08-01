package usecase

import (
	"errors"
	"fmt"

	client "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/client/interface"
	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/domain"
	interfaces "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/repository/interface"
	services "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/usecase/interface"
	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/utils/models"
)

type cartUseCase struct {
	cartRepository interfaces.CartRepository
	productClient client.ProductClient
}

func NewCartUseCase(repo interfaces.CartRepository,productClient client.ProductClient) services.CartUseCase {

	return &cartUseCase{
		cartRepository: repo,
		productClient: productClient,
	}
}



func (c *cartUseCase) AddToCart(productID int, userID int) (int,error) {


	productExist, err := c.productClient.DoesProductExist(productID)
	if err != nil {
		return 400,err
	}

	if !productExist {
		return 400,errors.New("product does not exists")
	}

	productPrice,err := c.productClient.GetProductPriceFromID(productID)
	if err != nil {
		return 400,err
	}

	productExistInCart, err := c.cartRepository.DoesProductExistInCart(productID)
	if err != nil {
		return 400,err
	}

	if !productExistInCart {
		quantity := 1
		err := c.cartRepository.InsertNewProductInCart(productID,quantity,productPrice,userID)
		if err != nil {
			return 400,err
		}
	} else {
		err := c.cartRepository.IterateQuantityInCart(productID,userID,productPrice)
		if err != nil {
			return 400,err
		}
	}

	return 201,nil
}



func (c *cartUseCase)	DisplayCArt(userID int) (models.CartResponse,error) {

	userCarts, err := c.cartRepository.GetUserCartFromID(userID)

	_, _ = userCarts,err
	productMap := []int{}

	for _,p := range userCarts {
		productMap = append(productMap,int(p.ProductID))
	}

	fmt.Println("productMap : ",productMap)
	
	m, err := c.productClient.GetProductNamesFromID(productMap)
	if err != nil {
		return models.CartResponse{},err
	}

	var GrandTotal float64
	for i := range userCarts {
		GrandTotal += userCarts[i].TotalPrice
		productID := userCarts[i].ProductID
		if movieName, ok := m[int(productID)]; ok {
			userCarts[i].MovieName = movieName
		} 
	}

	
	return models.CartResponse{
		TotalPrice: GrandTotal,
		Cart: userCarts,
	},nil

}

func (c *cartUseCase) OrderItemsFromCarts(orderItems domain.Order) (int,error)	 {

	grandTotal := 0.0
	for _,val := range orderItems.Items {
		grandTotal += val.Price
	}

	orderItems.TotalPrice = grandTotal

	// TO DO 
	// if address exist
	// if the cart corresponds to the same user

	if orderItems.Items == nil {
		return 0,errors.New("cart empty")
	}

	orderID, err := c.cartRepository.CreateNewOrder(orderItems)
	if err != nil {
		return 0,err
	}

	err = c.cartRepository.DeleteCartItems(orderItems.UserID)
	if err != nil {
		return 0,err
	}	

	return orderID,nil
}

