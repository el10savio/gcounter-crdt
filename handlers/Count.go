package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Count is the HTTP handler used to return
// the total count the GCounter nodes in the cluster
func Count(w http.ResponseWriter, r *http.Request) {
	// Sync the GCounters if multiple nodes
	// are present in a cluster
	if len(GetPeerList()) != 0 {
		GCounter, _ = Sync(GCounter)
	}

	// Get the total count in the GCounter
	count := GCounter.GetTotal()

	// DEBUG log in the case of success
	// indicating the total count
	log.WithFields(log.Fields{
		"count": count,
	}).Debug("successful gcounter total")

	// json encode response count
	json.NewEncoder(w).Encode(count)
}
