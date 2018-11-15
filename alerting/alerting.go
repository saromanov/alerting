package alerting

import "github.com/saromanov/alerting/structs"

// Alerting defines main interface for sending alerts
type Alerting interface {
	New() error
	Send(*structs.Message) (*structs.MessageResponse, error)
}
