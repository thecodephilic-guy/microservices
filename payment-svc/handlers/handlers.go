package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/paymentintent"
	"github.com/thecodephilic-guy/payment-svc/models"
)

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

// ProcessPayment handles incoming payment requests
func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set Stripe API Key
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Create a PaymentIntent
	pi, err := paymentintent.New(&stripe.PaymentIntentParams{
		Amount:             stripe.Int64(int64(req.Amount * 100)), // Convert to cents
		Currency:           stripe.String("usd"),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		PaymentMethod:      stripe.String("pm_card_visa"),
		Confirm:            stripe.Bool(true), // Try to confirm the payment immediately
	})

	if err != nil {
		log.Println("Payment processing failed:", err)
		resp := models.PaymentResponse{
			OrderID:      req.ID,
			Status:       "failed",
			ErrorMessage: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Check PaymentIntent status
	paymentStatus := "pending" // Default is pending
	if pi.Status == stripe.PaymentIntentStatusSucceeded {
		paymentStatus = "completed"
	} else if pi.Status == stripe.PaymentIntentStatusRequiresAction {
		paymentStatus = "pending" // 3D Secure required
	} else if pi.Status == stripe.PaymentIntentStatusCanceled {
		paymentStatus = "failed"
	}

	resp := models.PaymentResponse{
		OrderID:   req.ID,
		Status:    paymentStatus,
		PaymentID: pi.ID,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
