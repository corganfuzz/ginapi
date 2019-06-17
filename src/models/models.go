package models

//Product Model
type Product struct {
	Image       string `json:"img"`
	ImageAlt    string `json:"imgalt"`
	Price       string `json:"price"`
	Promotion   string `json:"promotion"`
	ProductName string `json:"productname"`
	Description string `json:"desc"`
}

// Customer Model
type Customer struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	LoggedIn  bool   `json:"loggedin"`
}

// Order Model , this includes 2 models from the structs above
type Order struct {
	Product
	Customer
	CustomerID   int    `json:"customer_id"`
	ProductID    string `json:"product_id"`
	Price        string `json:"sell_price"`
	PurchaseDate string `json:"purchase_date"`
}
