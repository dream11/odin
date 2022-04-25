package backend

import (
	"encoding/json"
	"path"

	"github.com/dream11/odin/api/profile"
)

// Profile entity
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

// ListProfiles : list profiles
func (s *Profile) ListProfiles(profileName, serviceName string) ([]profile.List, error) {
	client := newApiClient()
	client.QueryParams["name"] = profileName
	client.QueryParams["service"] = serviceName

	response := client.action(profileEntity, "GET", nil)
	response.Process(true)

	var serviceResponse profile.ListResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *Profile) DescribeProfile(profileName string) (profile.Describe, error) {
	client := newApiClient()

	response := client.action(path.Join(profileEntity, profileName), "GET", nil)
	response.Process(true)

	var serviceResponse profile.DescribeResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}
