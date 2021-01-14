package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID int 	`json:"id"`
	Name string `json:"name"`
	Price float32 `json:"price"`
}

var Products []*Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there")
}

func handleGetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Products)
}

func handleGetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"]) 

	for _, product := range Products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
		}
	}	
}

func handlePostProduct(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	
	var product *Product
	json.Unmarshal(body, &product)
	product.ID = len(Products)+1

	Products = append(Products, product)

	json.NewEncoder(w).Encode(product)
}

func handlePutProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	
	for _, product := range Products {
		if product.ID == id {
			json.Unmarshal(body, &product)
			fmt.Print(product)
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"]) 
	
	for index, product := range Products {
		if product.ID == id {
			Products = append(Products[:index], Products[index+1:]...)
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}


func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/products", handleGetProducts).Methods("GET")
	router.HandleFunc("/products", handlePostProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handleGetProduct).Methods("GET")
	router.HandleFunc("/products/{id}", handlePutProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handleDeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	Products = []*Product{
		{ID: 1, Name: "table", Price: 1000},
		{ID: 2, Name: "lighter", Price: 20},
	}

	handleRequests()
}