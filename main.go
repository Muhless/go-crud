package main

import (
	"log"
	"net/http"

	"github.com/Muhless/go-crud/config"
	categoriescontroller "github.com/Muhless/go-crud/controllers/categoriesController"
	homecontroller "github.com/Muhless/go-crud/controllers/homeController"
)

func main() {
	config.ConnectDB()

	// homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories", categoriescontroller.Index)
	http.HandleFunc("/categories/add", categoriescontroller.Add)
	http.HandleFunc("/categories/edit", categoriescontroller.Edit)
	http.HandleFunc("/categories/delete", categoriescontroller.Delete)

	log.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
