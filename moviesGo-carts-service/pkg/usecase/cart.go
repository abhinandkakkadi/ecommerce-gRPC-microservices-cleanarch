package usecase

import (
	"fmt"

	client "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/client/interface"
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

	fmt.Println("it reached here")
	productExist, err := c.productClient.DoesProductExist(productID)
	if err != nil {
		return 400,err
	}

	fmt.Println("it did not reached here")

	fmt.Println("product exist result : ",productExist)
	return 200,nil
}



func (c *cartUseCase)	DisplayCArt(userID int) (models.CartResponse,error) {

	return models.CartResponse{},nil

}

