package email

import (
	"fmt"

	"github.com/go-gomail/gomail"
)

// Email defines sending of alerts via email
type Email struct {
}

func New() error {
	return nil
}

// Send provides sending of email
func Send() error {
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("unable to send email: %v", err)
	}
	return nil
}
