package structs

// Message defines struct for alert message
type Message struct {
	Author    string
	Text      string
	Timestamp uint64
	Case      string
	Level     uint
}
