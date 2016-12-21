package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/meatballhat/negroni-logrus"
	_ "github.com/lib/pq"
	log "github.com/Sirupsen/logrus"

	"github.com/lumengxi/mockingbird/mockingbird/models"
	"github.com/lumengxi/mockingbird/mockingbird"
)


var (
	db *sql.DB
	mockers = []models.Mocker{}
)

func getHome(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("Get some mockers!")
}

func makeMocker(w http.ResponseWriter, req *http.Request) {
	var mocker models.Mocker
	_ = json.NewDecoder(req.Body).Decode(&mocker)
	mocker.ID = uuid.NewV4().String()
	mockers = append(mockers, mocker)

	json.NewEncoder(w).Encode(mockers)
}


func getMocker(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range mockers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Mocker{})
}

func getMockers(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(mockers)
}

func init() {
	mockingbird.InitLogger()
	mockingbird.InitDatabase("postgres://postgres:postgres@localhost/mockingbird?sslmode=disable")
}

	func main() {
		router := mux.NewRouter()

		// Routes
	router.HandleFunc("/", getHome).Methods("GET")
	router.HandleFunc("/mockers", getMockers).Methods("GET")
	router.HandleFunc("/mockers", makeMocker).Methods("POST")
	router.HandleFunc("/mockers/{id}", getMocker).Methods("GET")

	n := negroni.New()
	n.Use(negronilogrus.NewCustomMiddleware(log.DebugLevel, &log.JSONFormatter{}, "mockingbird"))
	n.Use(negroni.NewRecovery())
	n.UseHandler(router)

	log.Info("Mockingbird server started!")
	http.ListenAndServe(":1234", n)
}
