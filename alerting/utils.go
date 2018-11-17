package alerting

import "github.com/satori/go.uuid"

// GenID provides generation of UUID
func GenID() string {
	return uuid.Must(uuid.NewV4())
}
