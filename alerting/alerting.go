package alerting

import (
	"errors"
	"fmt"

	"github.com/robfig/cron"
	"github.com/saromanov/alerting/storage"
	"github.com/saromanov/alerting/storage/bolt"
	"github.com/saromanov/alerting/structs"
)

var errNoDBInit = errors.New("db is not initialized")

// Provider defines main interface for sending alerts
type Provider interface {
	New() error
	Send(*structs.Message) (*structs.MessageResponse, error)
}

// App provides entry point for api
type App struct {
	store     storage.Storage
	providers []Provider
}

// New creates app
func New(c *Config) *App {
	if len(c.Providers) == 0 {
		panic("providers is not defined")
	}
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
	if a.store == nil {
		return errNoDBInit
	}
	m.ID = GenID()
	err := a.store.Set(m)
	if err != nil {
		return fmt.Errorf("unable to collect message: %v", err)
	}
	return nil
}

// Run startes instance of the app for handling of collected messages
func (a *App) Run() error {
	c := cron.New()
	c.AddFunc("@every 1h", func() {
		messages, err := a.store.View()
		if err != nil {
			return
		}
		for _, m := range messages {
			fmt.Println(m)
		}
	})
	c.Start()
	for {
	}
}
