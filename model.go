package main

import (
	"database/sql"
	"errors"
)

type product struct {
	ID int 		`json:"id"`
	Name string 		`json:"name"`
	Price float32 	`json:"price"`
}

func (p *product) createProduct(db *sql.DB) (*product, error) {
	return nil, errors.New("Not Implemented")
}

func (p *product) updateProduct(db *sql.DB) (*product, error) {
	return nil, errors.New("Not Implemented")
}

func (p *product) getProduct(db *sql.DB) (*product, error) {
	return nil, errors.New("Not Implemented")
}

func (p *product) getProducts(db *sql.DB, limit, offset int) ([]*product, error) {
	return nil, errors.New("Not Implemented")
}

func (p *product) deleteProduct(db *sql.DB) (*product, error) {
	return nil, errors.New("Not Implemented")
}

