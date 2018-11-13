package structs

// Message defines struct for alert message
type Message struct {
	ID        string
	Author    string
	Text      string
	Timestamp uint64
	Case      string
	Level     uint
}
