package hawk

import (
	"crypto/sha256"
	"github.com/tent/hawk-go"
	"net/http"
)

// Creates a client Auth
func getAuth(key string, secret string, http *http.Request) *hawk.Auth {
	credentials := &hawk.Credentials{
		ID:   key,
		Key:  secret,
		Hash: sha256.New,
	}
	return hawk.NewRequestAuth(http, credentials, 0)
}

// Get Hawk headers from the given key, secret and http request
func getHeaders(key string, secret string, http *http.Request) string {
	auth := getAuth(key, secret, http)
	return auth.RequestHeader()
}
func postHeaders(key string, secret string, http *http.Request, payload string) string {
	auth := getAuth(key, secret, http)
	h := auth.PayloadHash(payload)
	auth.SetHash(h)
	return auth.RequestHeader()
}

// GenerateHeaders generate "X-Request-Auth" (Hawk) headers
// method is the http request method type ("GET", "PUT")
// url is the url to request
// key is the client id for Hawk authentication
// secret is the access token for Hawk authentication
func GenerateHeaders(method string, url string, key string, secret string) (string, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	return getHeaders(key, secret, req), nil
}

// GenerateHeadersWithPayload generate "X-Request-Auth" (Hawk) headers with a payload
// method is the http request method type ("POST", "PUT")
// url is the url to request
// key is the key for Hawk authentication
// secret is the access token for Hawk authentication
// payload is the payload that will be sent
func GenerateHeadersWithPayload(method string, url string, key string, secret string, payload string) (string, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	return postHeaders(key, secret, req, payload), nil
}
