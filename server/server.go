package server

import (
	"log"
	"net/http"
	"github.com/go-chi/chi"
)

// metrics provides output for metrics
func metrics(w http.ResponseWriter, r *http.Request) {

}

// config provides output for config
func config(w http.ResponseWriter, r *http.Request) {

}


// Create provides initialization of the server
func Create(addr string) {
	r := chi.NewRouter()
	r.Get("/v1/metrics", metrics)
	r.Get("/v1/config", config)
	log.Fatal(http.ListenAndServe(addr, r))
}
