package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
	"net/http"
)


func init() {
	InitLogger()
	InitDatabase("postgres://postgres:postgres@localhost/mockingbird?sslmode=disable")
}

// curl -X POST -d '{"name":"test", "config":{"body":"hello world!"}}' localhost:1234/mocker

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", GetHome).Methods("GET")
	router.HandleFunc("/mockers", GetMockers).Methods("GET")
	router.HandleFunc("/mocker", MakeMocker).Methods("POST")
	router.HandleFunc("/mockers/{id}", GetMocker).Methods("GET")

	n := negroni.New()
	n.Use(negronilogrus.NewCustomMiddleware(log.DebugLevel, &log.JSONFormatter{}, "mockingbird"))
	n.Use(negroni.NewRecovery())
	n.UseHandler(router)

	log.Info("Starting Mockingbird server...")
	http.ListenAndServe(":1234", n)
}
