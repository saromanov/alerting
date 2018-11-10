package alerting

// Alerting defines main interface for sending alerts
type Alerting interface {
	New() error
	Send() error
}
