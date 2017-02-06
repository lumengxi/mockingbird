package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
	"net/http"
	"os"
)


func init() {
	InitLogger()
	InitDatabase("postgres://postgres:postgres@localhost/mockingbird?sslmode=disable")
}

// curl -X POST -d '{"name":"test", "config":{"body":"hello world!"}}' localhost:1234/mocker

func main() {

	router := mux.NewRouter()

	// Routes
	router.Handle("/", HandlerWithError{GetHomeHandler}).Methods("GET")
	router.Handle("/mockers", HandlerWithError{GetMockerConfigsHandler}).Methods("GET")
	router.Handle("/mocker", HandlerWithError{MakeMockerHandler}).Methods("POST")
	router.Handle("/mockers?id={id}", HandlerWithError{GetMockerConfigHandler}).Methods("GET")

	n := negroni.New()
	n.Use(negronilogrus.NewCustomMiddleware(log.DebugLevel, &log.JSONFormatter{}, "mockingbird"))
	n.Use(negroni.NewRecovery())
	n.UseHandler(router)

	log.Info("Starting Mockingbird server...")
	log.Fatal(http.ListenAndServe(":1234", n))
}
