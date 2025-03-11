package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
)

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

func SendEmail(toEmail string, amount float64, newStatus string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SMTP_EMAIL")
	senderPassword := os.Getenv("SMTP_PASSWORD")

	e := email.NewEmail()
	e.From = fmt.Sprintf("Notification Service <%s>", senderEmail)
	e.To = []string{toEmail}
	e.Subject = "Order Status Update"
	e.Text = []byte(fmt.Sprintf("Dear Customer,\n\nStatus of payment for amount %.2f is: %s.\n\nThank you for shopping with us!", amount, newStatus))

	// Authenticate and send email
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	err := e.Send(smtpHost+":"+smtpPort, auth)
	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	log.Println("Email sent successfully to:", toEmail)
	return nil
}
