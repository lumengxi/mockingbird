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

	env := &Env{
		Port: os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}

	router := mux.NewRouter()

	// Routes
	router.Handle("/", HandlerMixin{env, GetHomeHandler}).Methods("GET")
	router.Handle("/mockers", HandlerMixin{env, GetMockers}).Methods("GET")
	router.Handle("/mocker", HandlerMixin{env, MakeMockerHandler}).Methods("POST")
	router.Handle("/mockers?id={id}", HandlerMixin{env, GetMockerHandler}).Methods("GET")

	n := negroni.New()
	n.Use(negronilogrus.NewCustomMiddleware(log.DebugLevel, &log.JSONFormatter{}, "mockingbird"))
	n.Use(negroni.NewRecovery())
	n.UseHandler(router)

	log.Info("Starting Mockingbird server...")
	log.Fatal(http.ListenAndServe(":1234", n))
}
