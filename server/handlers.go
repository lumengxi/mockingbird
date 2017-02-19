package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	mockers = []Mocker{}
)

// GetHomeHandler renders a simple homepage
func GetHomeHandler(w http.ResponseWriter, req *http.Request) error {
	return json.NewEncoder(w).Encode("Get some mockers!")
}

// MakeMockerHandler creates an active `Mocker`
func MakeMockerHandler(w http.ResponseWriter, req *http.Request) error {
	var mocker Mocker
	_ = json.NewDecoder(req.Body).Decode(&mocker)
	mocker.ID = uuid.NewV4().String()
	mocker.Status = true
	mockers = append(mockers, mocker)

	return json.NewEncoder(w).Encode(mockers)
}

// GetMockerConfigHandler returns mocker config by id
func GetMockerConfigHandler(w http.ResponseWriter, req *http.Request) error {
	mockerID := req.URL.Query().Get(":id")

	for _, mocker := range mockers {
		if mocker.ID == mockerID {
			return json.NewEncoder(w).Encode(mocker)
		}
	}

	return StatusError{500, fmt.Errorf("Cannot find requested mockerID: %d", mockerID)}
}

// GetMockerConfigsHandler returns all mockers' configs
func GetMockerConfigsHandler(w http.ResponseWriter, req *http.Request) error {
	return json.NewEncoder(w).Encode(mockers)
}

// SetMockerStatusHandler modifies a mocker status
func SetMockerStatusHandler(w http.ResponseWriter, req *http.Request) error {
	mockerId := req.URL.Query().Get(":id")
	mockerStatusParam := req.URL.Query().Get("status")

	targetStatus, err := strconv.ParseBool(mockerStatusParam)
	if err != nil {
		return StatusError{500, fmt.Errorf("Cannot parse mocker status input to bool: %s", mockerStatusParam)}
	}

	for _, mocker := range mockers {
		if mocker.ID == mockerId {
			log.Info("Set mocker [%d] status from %s to %s", mocker.ID, mocker.Status, targetStatus)
			mocker.Status = targetStatus
		}
	}

	return json.NewEncoder(w).Encode(&Mocker{})
}

// makeMockerResponse creates a `http.Response` based on mocker config
func makeMockerResponse(mockerConfig MockerConfig) http.Response {
	return http.Response{
		Header:     mockerConfig.MakeHeaders(),
		StatusCode: mockerConfig.StatusCode,
		Body:       ioutil.NopCloser(bytes.NewBufferString(mockerConfig.Body)),
	}
}

// GetMockerHandler renders a mocker-configured HTTP response
func GetMockerHandler(w http.ResponseWriter, req *http.Request) error {
	mockerID := req.URL.Query().Get(":id")

	for _, mocker := range mockers {
		if mocker.ID == mockerID {
			resp := makeMockerResponse(mocker.MockerConfig)
			return resp.Write(w)
		}
	}

	return StatusError{500, fmt.Errorf("Cannot find mocker by Id: %d", mockerID)}
}

// HealthCheckHandler setups basic application healthcheck
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}
