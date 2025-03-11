package models

// Notification struct
type Notification struct {
	Email   string  `json:"email"`
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}

// NotificationResponse struct
type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	OrderID string `json:"order_id,omitempty"`
	Status  string `json:"status,omitempty"`
	Email   string `json:"email,omitempty"`
}
