package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/pat"
	_ "github.com/lib/pq"
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
	"net/http"
)


func init() {
	InitLogger()
	InitDatabase("postgres://postgres:postgres@localhost/mockingbird?sslmode=disable")
}


func main() {

	router := pat.New()

	// Routes
	router.Handle("/", HandlerWithError{GetHomeHandler}).Methods("GET")
	router.Handle("/create-mocker", HandlerWithError{MakeMockerHandler}).Methods("POST")
	router.Handle("/show-mocker-configs", HandlerWithError{GetMockerConfigsHandler}).Methods("GET")
	router.Handle("/show-mocker-config/{id}", HandlerWithError{GetMockerConfigHandler}).Methods("GET")
	router.Handle("/mocker/{id}", HandlerWithError{GetMockerHandler}).Methods("GET")

	n := negroni.New()
	n.Use(negronilogrus.NewCustomMiddleware(log.DebugLevel, &log.JSONFormatter{}, "mockingbird"))
	n.Use(negroni.NewRecovery())
	n.UseHandler(router)

	log.Info("Starting Mockingbird server...")
	log.Fatal(http.ListenAndServe(":1234", n))
}
