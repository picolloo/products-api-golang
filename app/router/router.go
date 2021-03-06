package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/picolloo/productdex/app"
	uuid "github.com/satori/go.uuid"
)


var Products = []*app.Product{
	app.NewProduct("table",1000),
	app.NewProduct("lighter",20),
}

func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/products", handleGetProducts).Methods("GET")
	router.HandleFunc("/products", handlePostProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handleGetProduct).Methods("GET")
	router.HandleFunc("/products/{id}", handlePutProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handleDeleteProduct).Methods("DELETE")


	return router
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there")
}

func handleGetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(Products)
}

func handleGetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.FromString(vars["id"])

	for _, product := range Products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
		}
	}	
}

func handlePostProduct(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	product := app.NewProduct(data["name"].(string), float32(data["price"].(float64)))

	Products = append(Products, product)

	json.NewEncoder(w).Encode(product)
}

func handlePutProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.FromString(vars["id"])
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	
	for _, product := range Products {
		if product.ID == id {
			json.Unmarshal(body, &product)
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := uuid.FromString(vars["id"])
	
	for index, product := range Products {
		if product.ID == id {
			Products = append(Products[:index], Products[index+1:]...)
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}