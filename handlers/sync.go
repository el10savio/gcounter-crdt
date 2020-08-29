package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"../gcounter"
)

// Sync merges multiple GCounters present in a network to get them in sync
// It does so by obtaining the GCounter from each node in the cluster
// and performs a merge operation with the local GCounter
func Sync(GCounter gcounter.GCounter) (gcounter.GCounter, error) {
	// Obtain addresses of peer nodes in the cluster
	peers := GetPeerList()

	// Return the local GCounter back if no peers
	// are present along with an error
	if len(peers) == 0 {
		return GCounter, errors.New("nil peers present")
	}

	// Iterate over the peer list and send a /gcounter/values GET request
	// to each peer to obtain its GCounter
	for _, peer := range peers {
		peerGCounter, err := SendListRequest(peer)
		if err != nil {
			log.WithFields(log.Fields{"error": err, "peer": peer}).Error("failed sending gcounter values request")
			continue
		}

		// Skip merge if the peer's GCounter is empty
		if len(peerGCounter.Count) == 0 {
			continue
		}

		// Merge the peer's GCounter with our local GCounter
		GCounter = gcounter.Merge(GCounter, peerGCounter)
	}

	// DEBUG log in the case of success
	// indicating the new GCounter
	log.WithFields(log.Fields{
		"count": GCounter.Count,
	}).Debug("successful gcounter sync")

	// Return the synced new GCounter
	return GCounter, nil
}

// SendListRequest is used to send a GET /gcounter/values
// to peer nodes in the cluster
func SendListRequest(peer string) (gcounter.GCounter, error) {
	var _gcounter gcounter.GCounter

	// Return an empty GCounter followed by an error if the peer is nil
	if peer == "" {
		return _gcounter, errors.New("empty peer provided")
	}

	// Resolve the Peer ID and network to generate the request URL
	url := fmt.Sprintf("http://%s.%s/gcounter/values", peer, GetNetwork())
	response, err := SendRequest(url)
	if err != nil {
		return _gcounter, err
	}

	// Return an empty GCounter followed by an error
	// if the peer's response is not HTTP 200 OK
	if response.StatusCode != http.StatusOK {
		return _gcounter, errors.New("received invalid http response status:" + string(response.StatusCode))
	}

	// Decode the peer's GCounter to be usable by our local GCounter
	var values map[string]int
	err = json.NewDecoder(response.Body).Decode(&values)
	if err != nil {
		return _gcounter, err
	}

	// Return the decoded peer's GCounter
	_gcounter.Count = values
	return _gcounter, nil
}
