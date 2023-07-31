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

type ProductsReceiver struct {
	MovieName           string  `json:"movie_name binding:required"`
	GenreID             uint    `json:"genre_id binding:required"`
	ReleaseYear         string  `json:"release_year binding:required"`
	Format              string  `json:"format binding:required"`
	Director            string  `json:"director binding:required"`
	ProductsDescription string  `json:"products_description binding:required"`
	Runtime             float64 `json:"run_time binding:required"`
	Language            string  `json:"language binding:required"`
	StudioID            uint    `json:"studio_id binding:required"`
	Quantity            int     `json:"quantity binding:required"`
	Price               float64 `json:"price binding:required"`
}

type ProductName struct {
	ID        int
	MovieName string
}
