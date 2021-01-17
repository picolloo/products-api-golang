package app

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID uuid.UUID 					`json:"id"`
	Name string 		`json:"name"`
	Price float32 	`json:"price"`
}



func NewProduct(name string, price float32) *Product {
	return &Product{
		ID: uuid.NewV4(),
		Name: name,
		Price: price,
	}
}
