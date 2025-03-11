package models

import "time"

type PaymentRequest struct {
	ID        string    `json:"id,omitempty"` // UUID as Primary Key
	UserID    string    `json:"user_id"`      // Indexed for faster lookup
	Item      string    `json:"item"`
	Amount    int       `json:"amount"`     // Ensuring positive values
	Status    string    `json:"status"`     // Default status is 'pending'
	CreatedAt time.Time `json:"created_at"` // Auto timestamp
	UpdatedAt time.Time `json:"updated_at"`
}

type PaymentResponse struct {
	OrderID      string `json:"order_id"`
	Status       string `json:"status"`
	PaymentID    string `json:"payment_id,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}
