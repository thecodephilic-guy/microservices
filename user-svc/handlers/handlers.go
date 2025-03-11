package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thecodephilic-guy/user-svc/config"
	"github.com/thecodephilic-guy/user-svc/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//this function will extract the name and email from body and create a new user into the database:
	//creating new user:
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Invalid request"}`, http.StatusBadRequest)
		return
	}

	//inserting into the database:
	result := config.DB.Create(&user)

	if result.Error != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, result.Error.Error()), http.StatusBadRequest)
		return
	}

	response := models.Response{
		Message:     "User Created",
		Explanation: "A new user with given information has been created",
		Data:        user,
	}
	json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//this function will respond with user in the database based on the id:

	// paramer extraction
	params := mux.Vars(r)

	//converting string into uint:
	num, _ := strconv.ParseUint(params["user_id"], 10, 64)
	// Convert to uint (since ParseUint returns uint64)
	var user_id uint = uint(num)

	//creating an instance of struct to hold the data:
	var user models.User

	//querying the database:
	result := config.DB.First(&user, user_id)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, result.Error.Error()), http.StatusBadRequest)
		return
	}

	respone := models.Response{
		Message:     "User Found",
		Explanation: "User with the given id is found",
		Data:        user,
	}
	json.NewEncoder(w).Encode(respone)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//a slice of custom type to store the response from the database:
	var users []models.User

	//querying the database:
	result := config.DB.Find(&users) //this is equivalent to select * from users
	if result.Error != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, result.Error.Error()), http.StatusBadRequest)
		return
	}

	response := models.Response{
		Message:     "Users Data",
		Explanation: "Data of all the users",
		Data:        users,
	}

	json.NewEncoder(w).Encode(response)
}
