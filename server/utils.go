package server

import (
	log "github.com/Sirupsen/logrus"
)

// PanicOnError panic on err
func PanicOnError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
