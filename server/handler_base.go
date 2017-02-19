package main

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

type HandlerWithError struct {
	H func(w http.ResponseWriter, req *http.Request) error
}

func (h HandlerWithError) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := h.H(w, req)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Info("Http error: %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			log.Info("Unknown Http error: 500 - %s", e)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
	}

}
