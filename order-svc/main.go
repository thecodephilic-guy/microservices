package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thecodephilic-guy/order-svc/config"
	"github.com/thecodephilic-guy/order-svc/handlers"
)

func main() {
	//connect to the database
	config.ConnectDB()

	// uncomment the below line to perform the migrations whenever an update in schema is needed
	// config.MigrateDB()

	//defining the routes of the API:
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		message := map[string]string{"message": "Order Service is working fine!"}
		json.NewEncoder(w).Encode(message)
	}).Methods("GET")

	router.HandleFunc("/api/orders", handlers.GetOrders).Methods("GET")
	router.HandleFunc("/api/order/{user_id}", handlers.GetOrder).Methods("GET")
	router.HandleFunc("/api/createorder", handlers.CreateOrder).Methods("POST")

	fmt.Println("User Service is running at http://localhost:8002")
	log.Fatal(http.ListenAndServe(":8002", router))
}
