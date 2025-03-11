package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thecodephilic-guy/user-svc/config"
	"github.com/thecodephilic-guy/user-svc/handlers"
)

func main() {
	//connect to the database
	config.ConnectDB()

	// uncomment the below line to perform the migrations whenever an update in schema is needed
	// config.MigrateDB()

	//create API routes
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		message := map[string]string{"message": "User Service is working fine!"}
		json.NewEncoder(w).Encode(message)
	}).Methods("GET")

	//routes:
	router.HandleFunc("/api/createuser", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/getuser/{user_id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/getusers", handlers.GetUsers).Methods("GET")

	fmt.Println("User Service is running at http://localhost:8001")
	log.Fatal(http.ListenAndServe(":8001", router))
}
