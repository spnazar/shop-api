package main

import (
	"fmt"
	"net/http"

	"shop-api/db"
	"shop-api/handlers"
)

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

func main() {
	database := db.Connect()
	if database == nil {
		return
	}
	defer database.Close()

	db.Init(database)

	productHandler := &handlers.ProductHandler{DB: database}
	categoryHandler := &handlers.CategoryHandler{DB: database}

	http.HandleFunc("/products", logger(productHandler.GetAll))
	http.HandleFunc("/products/create", logger(productHandler.Create))
	http.HandleFunc("/products/delete", logger(productHandler.Delete))
	http.HandleFunc("/categories", logger(categoryHandler.GetAll))

	fmt.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe(":8080", nil)
}
