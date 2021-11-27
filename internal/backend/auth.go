package backend

import (
	"encoding/json"

	"github.com/dream11/odin/api/auth"
)

// Auth entity
type Auth struct{}

// GetToken : get access & refresh tokens
func (a *Auth) GetToken(accessKey, secretAccessKey string) (auth.Auth, error) {
	client := newClient()

	reqBody := map[string]string{
		"client_id":     accessKey,
		"client_secret": secretAccessKey,
	}

	// TODO: validate if reqBody needs a json conversion or not
	// if not then remove json conversion at all backend calls
	response := client.action("auth", "POST", reqBody)
	response.Process(true) // process response and exit if error

	var authResponse auth.Auth
	err := json.Unmarshal(response.Body, &authResponse)

	return authResponse, err
}

// RefreshToken : refresh the tokens
func (a *Auth) RefreshToken(refreshToken string) (auth.Auth, error) {
	client := newClient()

	reqBody := map[string]string{
		"refresh_token": refreshToken,
	}

	// TODO: validate if reqBody needs a json conversion or not
	// if not then remove json conversion at all backend calls
	response := client.action("refreshtoken", "POST", reqBody)
	response.Process(true) // process response and exit if error

	var authResponse auth.Auth
	err := json.Unmarshal(response.Body, &authResponse)

	return authResponse, err
}
