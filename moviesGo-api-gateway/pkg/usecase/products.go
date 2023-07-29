package usecase

import (
	"errors"
	"fmt"

	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
)

type productUseCase struct {
	productRepo interfaces.ProductRepository
	cartRepo    interfaces.CartRepository
}

func NewProductUseCase(repo interfaces.ProductRepository, cartRepo interfaces.CartRepository) services.ProductUseCase {
	return &productUseCase{
		productRepo: repo,
		cartRepo:    cartRepo,
	}
}

func (pr *productUseCase) ShowAllProducts(page int, count int) ([]models.ProductOfferBriefResponse, error) {

	productsBrief, err := pr.productRepo.ShowAllProducts(page, count)
	// here memory address of each item in productBrief is taken so that a copy of each instance is not made while updating
	for i := range productsBrief {
		p := &productsBrief[i]
		if p.Quantity == 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}
	var combinedProductsOffer []models.ProductOfferBriefResponse
	for _, p := range productsBrief {
		var productOffer models.ProductOfferBriefResponse

		productOffer.ProductsBrief = p
		combinedProductsOffer = append(combinedProductsOffer, productOffer)
	}

	return combinedProductsOffer, err

}


func (pr *productUseCase) ShowAllProductsToAdmin(page int, count int) ([]models.ProductsBrief, error) {

	productsBrief, err := pr.productRepo.ShowAllProducts(page, count)
	if err != nil {
		return []models.ProductsBrief{}, err
	}
	fmt.Println(productsBrief)
	// here memory address of each item in productBrief is taken so that a copy of each instance is not made while updating
	for i := range productsBrief {
		p := &productsBrief[i]
		if p.Quantity == 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}

	return productsBrief, nil
}


func (pr *productUseCase) ShowIndividualProducts(id string) (models.ProductOfferLongResponse, error) {

	product, err := pr.productRepo.ShowIndividualProducts(id)
	if err != nil {
		return models.ProductOfferLongResponse{}, err
	}
	if product.MovieName == "" {
		return models.ProductOfferLongResponse{}, errors.New("record not available")
	}

	var productOfferResponse models.ProductOfferLongResponse

	productOfferResponse.ProductsResponse = product

	return productOfferResponse, nil

}

func (pr *productUseCase) AddProduct(product models.ProductsReceiver) (models.ProductResponse, error) {
	// this logic is to add the quantity of product if admin try to add duplicate product (have to work on this in the future)
	// alreadyPresent,err := cr.productRepo.CheckIfAlreadyPresent(c,product)

	// if err != nil {
	// 	return err
	// }

	// if alreadyPresent {
	// 	fmt.Println("it came here")
	// 	err := cr.productRepo.UpdateQuantity(c,product)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil
	// }

	productResponse, err := pr.productRepo.AddProduct(product)

	if err != nil {
		return models.ProductResponse{}, err
	}

	return productResponse, nil

}

