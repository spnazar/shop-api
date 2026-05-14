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

	http.HandleFunc("POST /login", handlers.Login)
	http.HandleFunc("GET /products", logger(handlers.AuthMiddleware(productHandler.GetAll)))
	http.HandleFunc("GET /products/{id}", logger(handlers.AuthMiddleware(productHandler.GetOne)))
	http.HandleFunc("POST /products", logger(handlers.AuthMiddleware(productHandler.Create)))
	http.HandleFunc("PUT /products/{id}", logger(handlers.AuthMiddleware(productHandler.Update)))
	http.HandleFunc("DELETE /products/{id}", logger(handlers.AuthMiddleware(productHandler.Delete)))
	http.HandleFunc("GET /categories", logger(categoryHandler.GetAll))

	fmt.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe(":8080", nil)
}
