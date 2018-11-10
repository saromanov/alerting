package email

import (
	"fmt"

	"github.com/go-gomail/gomail"
)

// Email defines sending of alerts via email
type Email struct {
	email string
}

func New() error {
	return nil
}

// Send provides sending of email
func (e *Email) Send() error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.email)
	m.SetHeader("To", e.email)
	m.SetHeader("Subject", "Alert!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("unable to send email: %v", err)
	}
	return nil
}
