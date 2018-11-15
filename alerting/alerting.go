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
