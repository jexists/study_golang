package models

import "time"

// 데이터 모델
// - 데이터 베이스에 저장된 데이터를 표현하는 자료구조

// 구조체 태그(struct tag)
// json문서에서 어떤 키에 해당하는지 나타냄
type Product struct {
	Image       string  `json:"img"`
	ImagAlt     string  `json:"imag_alt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
}

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_mame"`
	Email     string `json:"email"`
	LoggedIn  bool   `json:"logged_in"`
	Password  string `json:"password"`
}

type Order struct {
	Product
	Customer
	CustomerID   int       `json:"customer_id"`
	ProductID    int       `json:"product_id"`
	Price        float64   `json:"sell_price"`
	PurchaseDate time.Time `json:"purchase_date"`
}
