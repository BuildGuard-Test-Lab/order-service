package main

import (
	"log"
	"net/http"

	"github.com/BuildGuard-Test-Lab/order-service/internal/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/api/v1/orders", handlers.ListOrders)
	log.Println("Starting order-service on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
