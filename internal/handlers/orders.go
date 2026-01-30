package handlers

import (
	"fmt"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"status":"healthy","service":"order-service"}`)
}

func ListOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"orders":[],"total":0}`)
}
