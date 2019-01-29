package server

import (
	"log"
	"net/http"
)

// Create provides initialization of the server
func Create(addr string) {
	log.Fatal(http.ListenAndServe(addr, nil))
}
