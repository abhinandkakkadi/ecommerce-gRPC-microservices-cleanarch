package domain

type Products struct {
	ID                  uint        `json:"id" gorm:"unique;not null"`
	MovieName           string      `json:"movie_name"`
	SKU                 string      `json:"sku"`
	GenreID             uint        `json:"genre_id"`
	Genre               Genre       `json:"-" gorm:"foreignkey:GenreID;constraint:OnDelete:CASCADE"`
	Language            string      `json:"language"`
	Director            string      `json:"director"`
	ReleaseYear         string      `json:"release_year"`
	Format              string      `json:"format"`
	ProductsDescription string      `json:"products_discription"`
	RunTime             float64     `json:"run_time"`
	StudioID            uint        `json:"studio_id"`
	MovieStudio         MovieStudio `json:"-" gorm:"foreignkey:StudioID;constraint:OnDelete:CASCADE"`
	Quantity            int         `json:"quantity"`
	Price               float64     `json:"price"`
	Delete              bool        `json:"delete" gorm:"default:false"`
}

type Genre struct {
	ID        uint   `json:"id" gorm:"unique; not null"`
	GenreName string `json:"genre_name"`
}

type MovieStudio struct {
	ID     uint   `json:"id" gorm:"unique; not null"`
	Studio string `json:"studio"`
}
