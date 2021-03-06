package server

import (
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"time"
)

// InitLogger configures logger
func InitLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}

// InitDatabase starts database with healthcheck
func InitDatabase(address string) (*sql.DB, error) {
	db, err := sql.Open("postgres", address)
	if err != nil {
		return nil, fmt.Errorf("Database connection failure: %s", err.Error())
	}

	if err = db.Ping(); err != nil {
		cancel := time.NewTimer(5 * time.Second)
		attempts := 1

	PingLoop:
		for {
			select {
			case <-time.After(1 * time.Second):
				if err := db.Ping(); err != nil {
					attempts++
					continue PingLoop
				}
				break PingLoop

			case <-cancel.C:
				return nil, fmt.Errorf("Database connection failed after %d attempts", attempts)
			}
		}
	}

	log.Info("Database connection established: %s", address)

	return db, nil
}
