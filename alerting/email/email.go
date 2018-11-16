package email

import (
	"fmt"

	"github.com/go-gomail/gomail"
	"github.com/saromanov/alerting/structs"
)

// Email defines sending of alerts via email
type Email struct {
	email    string
	host     string
	port     int
	username string
	password string
}

// New creates a new email init
func New() error {
	return nil
}

// Send provides sending of email
func (e *Email) Send(m *structs.Message) (*structs.MessageResponse, error) {
	g := gomail.NewMessage()
	g.SetHeader("From", e.email)
	g.SetHeader("To", e.email)
	g.SetHeader("Subject", "Alert!")
	g.SetBody("text/html", m.Text)

	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(g); err != nil {
		return nil, fmt.Errorf("unable to send email: %v", err)
	}
	return &structs.MessageResponse{}, nil
}
