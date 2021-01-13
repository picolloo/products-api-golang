package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	ID int 	`json:"id"`
	Name string `json:"name"`
	Price float32 `json:"Price"`
}

var Products []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there")
}

func handleGetProduct(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Products)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/products", handleGetProduct)
	log.Fatal(http.ListenAndServe(":3000", nil,))
}

func main() {
	Products = []Product{
		{ID: 1, Name: "table", Price: 1000},
		{ID: 2, Name: "lighter", Price: 20},
	}

	handleRequests()
}