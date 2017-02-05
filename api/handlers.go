package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

var (
	mockers = []Mocker{}
)

func GetHome(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("Get some mockers!")
}

func MakeMocker(w http.ResponseWriter, req *http.Request) {
	var mocker Mocker
	_ = json.NewDecoder(req.Body).Decode(&mocker)
	mocker.ID = uuid.NewV4().String()
	mocker.Status = true
	mockers = append(mockers, mocker)

	json.NewEncoder(w).Encode(mockers)
}


func GetMocker(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range mockers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Mocker{})
}

func GetMockers(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(mockers)
}

func DisableMocker(w http.ResponseWriter, req *http.Request) {

}