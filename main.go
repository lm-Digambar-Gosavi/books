package main

import (
	"books/config"
	"books/handlers"
	"books/repository"
	"books/routes"
	"books/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// Dependency Injection (Repository -> Service -> Handler)
	bookRepo := repository.NewBookRepository(config.DB)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	r := routes.SetupRoutes(bookHandler)

	fmt.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
