package alerting

import (
	"github.com/saromanov/alerting/db"
	"github.com/saromanov/alerting/structs"
)

// Alerting defines main interface for sending alerts
type Alerting interface {
	New() error
	Send(*structs.Message) (*structs.MessageResponse, error)
}

// App provides entry point for api
type App struct {
	d *db.DB
}

// New creates app
func New(c *Config) *App {
	return &App{}
}

// Send provides sending of message
func (a *App) Send(m *structs.Message) error {
	return nil
}

// Collect provides collecting of allerts
// and its send notification after some time
func (a *App) Collect(m *structs.Message) error {
	return nil
}
