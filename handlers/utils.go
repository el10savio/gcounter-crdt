package handlers

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"
)

// GetPeerList Obtains Peer List
// From Environment Variable
func GetPeerList() []string {
	if os.Getenv("PEERS") == "" {
		return []string{}
	}
	return strings.Split(os.Getenv("PEERS"), ",")
}

// GetNetwork Obtains Network
// From Environment Variable
func GetNetwork() string {
	return os.Getenv("NETWORK") + ":8080"
}

// GetMyNodeIP Obtains The Node's IP
// From Environment Variable
func GetMyNodeIP() string {
	return os.Getenv("MY_NODE")
}

// SendRequest handles sending of an HTTP GET Request
func SendRequest(url string) (http.Response, error) {
	if url == "" {
		return http.Response{}, errors.New("empty url provided")
	}

	client := http.Client{
		Timeout: time.Duration(5 * 60 * time.Second),
	}

	response, err := client.Get(url)
	if err != nil {
		return http.Response{}, err
	}

	return *response, nil
}
