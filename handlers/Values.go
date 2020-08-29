package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Values is the HTTP handler to return the local GCounter's values
// without syncing it with other nodes in a cluster
func Values(w http.ResponseWriter, r *http.Request) {
	// Get the local GCounter values
	counter := GCounter.GetCount()

	// DEBUG log in the case of successfull
	// list indicating the counter
	log.WithFields(log.Fields{
		"counter": counter,
	}).Debug("successful gcounter values")

	// json encode response value
	json.NewEncoder(w).Encode(counter)
}
