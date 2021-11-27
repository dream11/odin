package backend

import (
	"path"
)

// Profile entity
type Profile struct{}

// root entity
var profileEntity = "profiles"

// CreateProfile : register a profile version with backend
func (p *Profile) CreateProfile(profile []byte) {
	client := newApiClient()

	response := client.action(profileEntity, "POST", profile)
	response.Process(true) // process response and exit if error
	// no return required
}

// DescribeProfile : describe a profile version or all versions of a profile
func (p *Profile) DescribeProfile(profile, version string) {
	client := newApiClient()
	client.QueryParams["version"] = version

	response := client.action(path.Join(profileEntity, profile), "GET", nil)
	response.Process(true) // process response and exit if error

	// TODO: parse response.Body into required structure and return
}

// ListProfiles : list profiles per team and describe versions
func (p *Profile) ListProfiles(team, version string) {
	client := newApiClient()
	client.QueryParams["team"] = team
	client.QueryParams["version"] = version

	response := client.action(profileEntity, "GET", nil)
	response.Process(true) // process response and exit if error

	// TODO: parse response.Body into required structure and return
}

// DeleteProfile : delete a profile version
func (p *Profile) DeleteProfile(profile, version string) {
	client := newApiClient()

	response := client.action(path.Join(profileEntity, profile, "version", version), "DELETE", nil)
	response.Process(true) // process response and exit if error
	// no return required
}
