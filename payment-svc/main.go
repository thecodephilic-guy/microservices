package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thecodephilic-guy/payment-svc/handlers"
)

func main() {
	//defining the routes of the API:
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		message := map[string]string{"message": "Payment Service is working fine!"}
		json.NewEncoder(w).Encode(message)
	}).Methods("GET")

	router.HandleFunc("/api/processpayment", handlers.ProcessPayment).Methods("POST")

	fmt.Println("User Service is running at http://localhost:8003")
	log.Fatal(http.ListenAndServe(":8003", router))
}
