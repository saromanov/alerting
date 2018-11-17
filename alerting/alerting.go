package alerting

import (
	"errors"
	"fmt"

	"github.com/saromanov/alerting/db"
	"github.com/saromanov/alerting/db/bolt"
	"github.com/saromanov/alerting/structs"
)

var errNoDBInit = errors.New("db is not initialized")

// Alerting defines main interface for sending alerts
type Alerting interface {
	New() error
	Send(*structs.Message) (*structs.MessageResponse, error)
}

// App provides entry point for api
type App struct {
	store db.DB
}

// New creates app
func New(c *Config) *App {
	return &App{
		store: bolt.New(),
	}
}

// Send provides sending of message
func (a *App) Send(m *structs.Message) error {
	return nil
}

// Collect provides collecting of allerts
// and its send notification after some time
func (a *App) Collect(m *structs.Message) error {
	if a.d == nil {
		return errNoDBInit
	}
	err := a.d.Set(m)
	if err != nil {
		return fmt.Errorf("unable to collect message: %v", err)
	}
	return nil
}
