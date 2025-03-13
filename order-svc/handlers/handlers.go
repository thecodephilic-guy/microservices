package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thecodephilic-guy/order-svc/config"
	"github.com/thecodephilic-guy/order-svc/models"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//creating a slice of custom type Order to hold the orders:
	var orders []models.Order

	//querying the database to get all the orders:
	result := config.DB.Find(&orders)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, result.Error.Error()), http.StatusBadRequest)
		return
	}

	//this handler will respond with all the orders in the database:
	response := models.Response{
		Message:     "All Orders",
		Explanation: "Detailed record of all the orders",
		Data:        orders,
	}

	json.NewEncoder(w).Encode(response)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, `{"error": "Invalid request"}`, http.StatusBadRequest)
		return
	}

	//setting the status of the order to pending:
	order.Status = "pending"

	// Now make a call to payment service to process the payment for this order based on order ID which is running on localhost 8003
	paymentServiceURL := "http://localhost:8003/api/processpayment"
	paymentReqBody, _ := json.Marshal(order) // Converting the order struct to json

	paymentResp, err := http.Post(paymentServiceURL, "application/json", bytes.NewBuffer(paymentReqBody))
	if err != nil {
		http.Error(w, `{"error": "Failed to contact payment service"}`, http.StatusInternalServerError)
		return
	}
	defer paymentResp.Body.Close()

	// Read response from payment service:
	var paymentResponse models.PaymentResponse
	if err := json.NewDecoder(paymentResp.Body).Decode(&paymentResponse); err != nil {
		http.Error(w, `{"error": "Invalid response from payment service"}`, http.StatusInternalServerError)
		return
	}

	// Check the payment status and respond accordingly
	switch paymentResponse.Status {
	case "completed":
		order.Status = "completed"
	case "pending":
		order.Status = "pending"
	case "failed":
		http.Error(w, `{"message": "Payment failed"}`, http.StatusBadRequest)
		return
	default:
		http.Error(w, `{"message": "Unknown payment status"}`, http.StatusInternalServerError)
		return
	}

	// Insert this order into the database:
	result := config.DB.Create(&order)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, result.Error.Error()), http.StatusBadRequest)
		return
	}

	// Fetch user data to send notification
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8001/api/getuser/%s", order.UserID), nil)
	if err != nil {
		http.Error(w, `{"error": "Failed to create request"}`, http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch user data"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, `{"error": "Failed to fetch user data"}`, resp.StatusCode)
		return
	}

	var userResponse models.UserResponse

	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		http.Error(w, `{"error": "Failed to decode user data"}`, http.StatusInternalServerError)
		return
	}

	// Prepare the notification payload
	notificationPayload := models.NotificationPayload{
		Email:   userResponse.Data.Email,
		OrderID: order.ID,
		Amount:  float64(order.Amount),
		Status:  order.Status,
	}

	payloadBytes, err := json.Marshal(notificationPayload)
	if err != nil {
		http.Error(w, `{"error": "Failed to encode notification payload"}`, http.StatusInternalServerError)
		return
	}

	// Send the notification request
	notifyReq, err := http.NewRequest("POST", "http://localhost:8004/api/notify", bytes.NewBuffer(payloadBytes))
	if err != nil {
		http.Error(w, `{"error": "Failed to create notification request"}`, http.StatusInternalServerError)
		return
	}
	notifyReq.Header.Set("Content-Type", "application/json")

	notifyResp, err := client.Do(notifyReq)
	if err != nil {
		http.Error(w, `{"error": "Failed to send notification request"}`, http.StatusInternalServerError)
		return
	}
	defer notifyResp.Body.Close()

	if notifyResp.StatusCode != http.StatusOK {
		http.Error(w, `{"error": "Failed to send notification"}`, notifyResp.StatusCode)
		return
	}

	response := models.Response{
		Message:     "New Order Placed",
		Explanation: "A new order is successfully placed",
		Data:        order,
	}
	json.NewEncoder(w).Encode(response)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user_id := params["user_id"]

	//creating a slice of struct to hold the data:
	var orders []models.Order

	//querying the database:
	result := config.DB.Where(&models.Order{UserID: user_id}).Find(&orders)

	//if no order is found with the given id:
	if result.RowsAffected == 0 {
		http.Error(w, `{"message" : "Order not found"}`, http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, result.Error.Error()), http.StatusBadRequest)
		return
	}

	response := models.Response{
		Message:     "Orders Found",
		Explanation: "Orders with the given user_id are found",
		Data:        orders,
	}
	json.NewEncoder(w).Encode(response)
}
