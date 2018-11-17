package db

import "github.com/saromanov/alerting/structs"

// DB defines
type DB interface {
	Set(m *structs.Message) error
	View() ([]*structs.Message, error)
	Get(id string) (*structs.Message, error)
}
