package main

import (
	"log"
	"net/http"

	"github.com/picolloo/productdex/app/router"
)



func handleRequests() {
	router := router.Router()

	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	handleRequests()
}