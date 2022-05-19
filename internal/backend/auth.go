package backend

import (
	"encoding/json"
	"fmt"

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

	response := client.actionWithRetry("secure/auth/", "POST", reqBody)
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

	response := client.actionWithRetry("secure/refreshtoken/", "POST", reqBody)
	response.Process(true) // process response and exit if error

	var authResponse []auth.Auth
	err := json.Unmarshal(response.Body, &authResponse)

	for _, token := range authResponse {
		if !token.Expired {
			return token, err
		}
	}

	return auth.Auth{}, fmt.Errorf("unable to find a valid active token")
}
