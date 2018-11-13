package structs

import "encoding/json"

// Message defines struct for alert message
type Message struct {
	ID        string
	Author    string
	Text      string
	Timestamp uint64
	Case      string
	Level     uint
	Namespace string
}

// Marshal provides marshaling of the message struct
func (m *Message) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
