package backend

import (
	"encoding/json"

	"github.com/dream11/odin/api/auth"
)

// Auth entity
type Auth struct{}

// GetToken : get access & refresh tokens
func (a *Auth) Authenticate(accessKey, secretAccessKey string) (string, error) {
	client := newClient()

	response := client.actionWithRetry("secure/authenticateuser/", "GET", nil)
	response.Process(true) // process response and exit if error

	var authResponse auth.Auth
	err := json.Unmarshal(response.Body, &authResponse)

	return authResponse.Response, err
}
