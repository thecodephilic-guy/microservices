package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thecodephilic-guy/notification-svc/models"
	"github.com/thecodephilic-guy/notification-svc/utils"
)

func NotifyUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//extracting request body and storing it in the custom struct:
	var req models.Notification

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request"}`, http.StatusBadRequest)
		return
	}

	// Call the email sending function
	err := utils.SendEmail(req.Email, req.Amount, req.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{Error: "Failed to send notification"})
		return
	}

	// Success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "Notification sent successfully",
		OrderID: req.OrderID,
		Status:  req.Status,
		Email:   req.Email,
	})

}
