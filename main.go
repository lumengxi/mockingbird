package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/pat"
	_ "github.com/lib/pq"
	"github.com/lumengxi/mockingbird/server"
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
	"net/http"
)


func init() {
	server.InitLogger()
	server.InitDatabase("postgres://postgres:postgres@localhost/mockingbird?sslmode=disable")
}


func main() {

	router := pat.New()

	// Routes
	router.Handle("/", server.HandlerWithError{server.GetHomeHandler}).Methods("GET")
	router.Handle("/create-mocker", server.HandlerWithError{server.MakeMockerHandler}).Methods("POST")
	router.Handle("/show-mocker-configs", server.HandlerWithError{server.GetMockerConfigsHandler}).Methods("GET")
	router.Handle("/show-mocker-config/{id}", server.HandlerWithError{server.GetMockerConfigHandler}).Methods("GET")
	router.Handle("/mocker/{id}", server.HandlerWithError{server.GetMockerHandler}).Methods("GET")

	n := negroni.New()
	n.Use(negronilogrus.NewCustomMiddleware(log.DebugLevel, &log.JSONFormatter{}, "mockingbird"))
	n.Use(negroni.NewRecovery())
	n.UseHandler(router)

	log.Info("Starting Mockingbird server...")
	log.Fatal(http.ListenAndServe(":1234", n))
}
