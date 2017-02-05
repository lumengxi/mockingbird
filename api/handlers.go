package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

var (
	mockers = []Mocker{}
)

func GetHome(w http.ResponseWriter, req *http.Request) error {
	json.NewEncoder(w).Encode("Get some mockers!")
}

func MakeMocker(w http.ResponseWriter, req *http.Request) error {
	var mocker Mocker
	_ = json.NewDecoder(req.Body).Decode(&mocker)
	mocker.ID = uuid.NewV4().String()
	mocker.Status = true
	mockers = append(mockers, mocker)

	return json.NewEncoder(w).Encode(mockers)
}


func GetMocker(w http.ResponseWriter, req *http.Request) error {
	mockerId := req.URL.Query().Get("id")

	for _, mocker := range mockers {
		if mocker.ID == mockerId {
			json.NewEncoder(w).Encode(mocker)
		} else {
			returnErr := fmt.Errorf("Cannot find requested mockerId: %d", mocker.ID)
			return StatusError{500, returnErr}
		}
	}
	return json.NewEncoder(w).Encode(&Mocker{})
}

func GetMockers(w http.ResponseWriter, req *http.Request) error {
	return json.NewEncoder(w).Encode(mockers)
}


func SetMockerStatus(w http.ResponseWriter, req *http.Request) error {
	mockerId := req.URL.Query().Get("id")
	mockerStatusParam := req.URL.Query().Get("status")

	targetStatus, err := strconv.ParseBool(mockerStatusParam)
	if err != nil {
		returnErr := fmt.Errorf("Cannot parse mocker status input to bool: %s", mockerStatusParam)
		log.Error(returnErr)
		return StatusError{500, returnErr}
	}

	for _, mocker := range mockers {
		if mocker.ID == mockerId {
			log.Info("Set mocker [%d] status from %s to %s", mocker.ID, mocker.Status, targetStatus)
			mocker.Status = targetStatus
		}
	}

	return json.NewEncoder(w).Encode(&Mocker{})
}
