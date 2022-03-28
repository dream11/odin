package backend

import (
	"encoding/json"
	"github.com/dream11/odin/api/profile"
)

// Service entity
type Profile struct{}

// root entity
var profileEntity = "profile"

// CreateProfile : register a profile with backend
func (s *Profile) CreateProfile(profileDefinition interface{}) (string, error) {
	client := newApiClient()

	response := client.action(profileEntity+"/", "POST", profileDefinition)
	response.Process(true) // process response and exit if error

	var serviceResponse profile.CreateResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response.Message, err
}
