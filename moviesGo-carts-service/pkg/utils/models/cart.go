package models


type Cart struct {
	ProductID  uint    
	MovieName  string  
	Quantity   float64 
	TotalPrice float64 
}

type CartResponse struct {
	UserName   string
	TotalPrice float64
	Cart       []Cart
}

