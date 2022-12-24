package main

import (
	"fmt"
	"golang_CRUD_rest_api_with_mongodb/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello!")
	r := router.Router()
	fmt.Println("Server is strating...")
	log.Fatal(http.ListenAndServe(":7000", r))

}
