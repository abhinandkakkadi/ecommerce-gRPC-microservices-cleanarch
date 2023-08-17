package repository

import (
	"fmt"

	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/domain"
	interfaces "github.com/abhinandkakkadi/moviesgo-carts-service/pkg/repository/interface"
	"github.com/abhinandkakkadi/moviesgo-carts-service/pkg/utils/models"
	"gorm.io/gorm"
)

type CartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &CartDatabase{
		DB: DB,
	}
}

func (c *CartDatabase) DoesProductExistInCart(productID int) (bool, error) {

	var productExist bool
	err := c.DB.Raw("SELECT EXISTS (SELECT 1 FROM carts WHERE product_id = ?) AS product_exist", productID).Scan(&productExist).Error
	if err != nil {
		return false, err
	}

	return productExist, nil

}

func (c *CartDatabase) InsertNewProductInCart(productID int, quantity int, productPrice float64, userID int) error {

	err := c.DB.Exec("INSERT INTO carts (user_id,product_id,quantity,total_price) values (?,?,?,?)", userID, productID, quantity, productPrice).Error
	if err != nil {
		return err
	}

	return nil

}

func (c *CartDatabase) IterateQuantityInCart(productID int, userID int, productPrice float64) error {

	err := c.DB.Exec("UPDATE CARTS SET total_price = total_price + ?,quantity = quantity + 1 WHERE user_id = ? and product_id = ?", productPrice, userID, productID).Error
	if err != nil {
		return err
	}

	return nil

}

func (c *CartDatabase) GetUserCartFromID(userID int) ([]models.Cart, error) {

	var carts []models.Cart
	err := c.DB.Raw("SELECT product_id, quantity, total_price FROM carts where user_id = ?", userID).Scan(&carts).Error
	if err != nil {
		return []models.Cart{}, err
	}

	fmt.Printf("%+v", carts)

	return carts, nil
}

func (c *CartDatabase) CreateNewOrder(orderProducts domain.Order) (int, error) {

	err := c.DB.Create(&orderProducts).Error
	if err != nil {
		return 0, err
	}

	return int(orderProducts.ID), nil

}

func (c *CartDatabase) DeleteCartItems(userID int) error {

	err := c.DB.Exec("DELETE FROM carts where user_id = ?", userID).Error
	if err != nil {
		return err
	}

	return nil

}
