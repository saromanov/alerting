package structs

import (
	"encoding/json"
	"fmt"
)

// Message defines struct for alert message
type Message struct {
	ID        string
	Author    string
	Text      string
	Timestamp uint64
	Case      string
	Level     uint
	Namespace string
	Subject string
}

// Marshal provides marshaling of the message struct
func (m *Message) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// String return text representation of Message
func (m *Message) String() string {
	return fmt.Sprintf("Author: %s\nText: %s\nCase: %s\n", m.Author, m.Text, m.Case)
}

// MessageResponse defines response after sending of message
// to provider
type MessageResponse struct {
	ChannelID string
	Timestamp int64
	ID string
	Provider string
}
