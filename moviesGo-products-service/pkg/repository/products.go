package repository

import (
	"fmt"

	interfaces "github.com/abhinandkakkadi/moviesgo-products-service/pkg/repository/interface"
	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/utils/models"
	"gorm.io/gorm"
)

type ProductDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductsRepository {
	return &ProductDatabase{
		DB: DB,
	}
}

func (p *ProductDatabase) ShowAllProducts(page int, count int) ([]models.ProductsBrief, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var productsBrief []models.ProductsBrief
	err := p.DB.Raw(`
		SELECT products.id, products.movie_name,products.sku,products.language,genres.genre_name AS genre, products.language,products.price,products.quantity
		FROM products
		JOIN genres ON products.genre_id = genres.id
		 limit ? offset ?
	`, count, offset).Scan(&productsBrief).Error

	if err != nil {
		return nil, err
	}

	return productsBrief, nil

}

func (p *ProductDatabase) AddProduct(productReceiver models.ProductsReceiver) (int, error) {

	var id int

	sku := productReceiver.MovieName + productReceiver.Format + productReceiver.Director
	err := p.DB.Raw("INSERT INTO products (movie_name, genre_id,language,director,release_year,format,products_description,run_time,studio_id,quantity,price,sku) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?) RETURNING id", productReceiver.MovieName, productReceiver.GenreID, productReceiver.Language, productReceiver.Director, productReceiver.ReleaseYear, productReceiver.Format, productReceiver.ProductsDescription, productReceiver.Runtime, productReceiver.StudioID, productReceiver.Quantity, productReceiver.Price, sku).Scan(&id).Error
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *ProductDatabase) GetGenreDetails() {

}

func (p *ProductDatabase) GetStudioDetails() {

}

func (p *ProductDatabase) DoesProductExist(productID int) (bool, error) {

	var productExist bool
	err := p.DB.Raw("SELECT EXISTS (SELECT 1 FROM products WHERE id = ?) AS user_exists", productID).Scan(&productExist).Error
	if err != nil {
		return false, err
	}
	fmt.Println("product status : ", productExist)
	return productExist, nil

}

func (p *ProductDatabase) GetProductPriceFromID(productID int) (float64, error) {

	var productPrice float64
	err := p.DB.Raw("SELECT price from products where id = ?", productID).Scan(&productPrice).Error
	if err != nil {
		return 0.0, err
	}

	return productPrice, nil

}

func (p *ProductDatabase) GetProductNameFromID(productIDS []int) ([]models.ProductName, error) {

	var productIDToName []models.ProductName

	err := p.DB.Raw("SELECT id, movie_name FROM products WHERE id = ANY($1)", productIDS).Scan(&productIDToName).Error
	if err != nil {
		return []models.ProductName{}, err
	}

	return productIDToName, nil

}
