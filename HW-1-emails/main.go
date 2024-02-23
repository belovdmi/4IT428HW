package main

import (
	"fmt"
)

// emailService reprezentuje službu pro odesílání a validaci e-mailů
type emailService struct{}


// Store e-mail
func (es *emailService) Store(email string) {
	fmt.Println("Stored email:", email)
}

// Send odesílá e-mail podle zadaného protokolu
func (es *emailService) Send(email string, protocol string) {
	switch protocol {
	case "SMTP":
		fmt.Println("Send email via SMTP:", email)
	case "IMAP":
		fmt.Println("Send email via IMAP:", email)
	case "POP3":
		fmt.Println("Send email via POP3:", email)
	default:
		fmt.Println("Send email via SMTP (default):", email)
	}

	es.Store(email)
}

// Validate ověřuje platnost e-mailové adresy
func (es *emailService) Validate(email string) {
	fmt.Println("Validate email:", email)
}

func main() {
	// Vytvoření instance emailService
	emailSvc := &emailService{}

	// Příklad použití
	email := "example@example.com"
	protocol := "SMTP"

	// Odeslání e-mailu
	emailSvc.Send(email, protocol)

	// Validace e-mailu
	emailSvc.Validate(email)
}