package storage

import "github.com/saromanov/alerting/structs"

// Storage defines interface for storage operations
type Storage interface {
	Set(m *structs.Message) error
	View() ([]*structs.Message, error)
	Get(id string) (*structs.Message, error)
}
