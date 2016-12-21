package mockingbird

import (
	log "github.com/Sirupsen/logrus"
)


func PanicOnError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
