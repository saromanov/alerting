package server

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/saromanov/alerting/alerting"
)

// metrics provides output for metrics
func metrics(w http.ResponseWriter, r *http.Request) {

}

// config provides output for config
func config(w http.ResponseWriter, r *http.Request) {

}


// Create provides initialization of the server
func Create(c *alerting.Config) error {
	r := chi.NewRouter()
	r.Get("/v1/metrics", metrics)
	r.Get("/v1/config", config)
	if err := http.ListenAndServe(c.ServerAddress, r); err != nil {
		return fmt.Errorf("unable to make server: %v", err)
	}
	return nil
}
