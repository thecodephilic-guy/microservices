package models

import (
	"time"
)

type Order struct {
	ID        string    `json:"id,omitempty" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // UUID as Primary Key
	UserID    string    `json:"user_id" gorm:"not null;index"`                                      // Indexed for faster lookup
	Item      string    `json:"item" gorm:"type:varchar(255);not null"`
	Amount    int       `json:"amount" gorm:"not null;check:amount >= 0"`                  // Ensuring positive values
	Status    string    `json:"status" gorm:"type:varchar(50);not null;default:'pending'"` // Default status is 'pending'
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`                          // Auto timestamp
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Response struct {
	Message     string      `json:"message"`
	Explanation string      `json:"explanation"`
	Data        interface{} `json:"data"`
}

type PaymentResponse struct {
	OrderID      string `json:"order_id"`
	Status       string `json:"status"`
	PaymentID    string `json:"payment_id,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

type UserResponse struct {
	Message     string `json:"message"`
	Explanation string `json:"explanation"`
	Data        struct {
		UserID int    `json:"user_id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
	} `json:"data"`
}

type NotificationPayload struct {
	Email   string  `json:"email"`
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}

var Orders = []Order{
	{ID: "101", UserID: "1", Item: "Laptop", Amount: 1500, Status: "Shipped"},
	{ID: "102", UserID: "2", Item: "Phone", Amount: 800, Status: "Under processing"},
}
