package backend

import (
	"encoding/json"
	"fmt"
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

func (s *Profile) DeleteProfile(profile string) {
	client := newApiClient()

	response := client.action(path.Join(profileEntity, profile)+"/", "DELETE", nil)
	response.Process(true)
}

func (s *Profile) DeployProfile(profileName, env, platform string, forceDeployServices []profile.ListEnvService, force bool) ([]profile.ProfileServiceDeploy, error) {
	client := newApiClient()
	client.QueryParams["env_name"] = env
	client.QueryParams["platform"] = platform
	client.QueryParams["force"] = fmt.Sprintf("%v", force)

	response := client.action(path.Join(profileEntity, "deploy", profileName, "env", env)+"/", "POST", forceDeployServices)
	response.Process(true)

	var serviceResponse profile.ProfileServiceDeployResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s Profile) ListEnvServices(profileName, env, filterBy string) ([]profile.ListEnvService, error) {
	client := newApiClient()
	client.QueryParams["filterBy"] = filterBy

	response := client.action(path.Join(profileEntity, profileName, "env", env, "service")+"/", "GET", nil)
	response.Process(true)

	var serviceResponse profile.ListEnvServiceResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *Profile) UndeployProfile(profileName, env string, forceDeployServices []profile.ListEnvService, force bool) ([]profile.ProfileServiceDeploy, error) {
	client := newApiClient()
	client.QueryParams["env_name"] = env
	client.QueryParams["force"] = fmt.Sprintf("%v", force)

	response := client.action(path.Join(profileEntity, "undeploy", profileName, "env", env)+"/", "DELETE", forceDeployServices)
	response.Process(true)

	var serviceResponse profile.ProfileServiceDeployResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (p *Profile) UpdateProfile(profileName string, profile interface{}) {
	client := newApiClient()

	response := client.action(path.Join(profileEntity, profileName)+"/", "PUT", profile)
	response.Process(true) // process response and exit if error
}
