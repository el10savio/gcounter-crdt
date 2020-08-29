package main

// The following implements the main Go
// package starting up the gcounter server

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"./handlers"
)

const (
	// PORT used for the GCounter server
	PORT = "8080"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := handlers.Router()

	log.WithFields(log.Fields{
		"port": PORT,
	}).Info("started GCounter node server")

	http.ListenAndServe(":"+PORT, r)
}
