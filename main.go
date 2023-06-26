package main

import (
	"fmt"
	"net/smtp"
)

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}
func main() {
	// Sender data.
	from := "mercymaina@infinitytechafrica.com"
	password := "you google app password."

	// Receiver email address.
	to := []string{
		"mercymaina567@gmail.com",
		"aliphonzanderitu@gmail.com",
	}

	message := []byte("This is a really unimaginative message, I know.")
	// Authentication.

	auth := smtp.PlainAuth(" ", from, password, "smtp.gmail.com")
	// Sending email.

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
	if err != nil {
		fmt.Printf("Errror sending email %s", err)
		return
	}
	fmt.Println("Email Sent!")
}
