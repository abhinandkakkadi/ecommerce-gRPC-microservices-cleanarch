package models

type Cart struct {
	ProductID  uint
	MovieName  string
	Quantity   float64
	TotalPrice float64
}

type CartResponse struct {
	TotalPrice float64
	Cart       []Cart
}

type OrderRequest struct {
	Cart      []Cart
	AddressID int
	UserId    int
}
