package models


type ProductsBrief struct {
	ID            int     `json:"id"`
	MovieName     string  `json:"movie_name"`
	Sku           string  `json:"sku"`
	Genre         string  `json:"genre"`
	Language      string  `json:"language"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
	ProductStatus string  `json:"product_status"`
}
