package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thecodephilic-guy/notification-svc/handlers"
)

func main() {
	//defining the routes of the API:
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		message := map[string]string{"message": "Notification Service is working fine!"}
		json.NewEncoder(w).Encode(message)
	}).Methods("GET")

	router.HandleFunc("/api/notify", handlers.NotifyUser).Methods("POST")

	fmt.Println("User Service is running at http://localhost:8004")
	log.Fatal(http.ListenAndServe(":8004", router))
}
