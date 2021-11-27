package backend

import (
	"path"
)

// Infra entity
type Infra struct{}

// root entity
var infraEntity = "infras"

// CreateInfra : create an empty infra
func (i *Infra) CreateInfra(infra string, infraDetails []byte) {
	client := newApiClient()

	response := client.action(path.Join(infraEntity, infra), "POST", infraDetails)
	response.Process(true) // process response and exit if error
	// no return required
}

// DescribeInfra : describe an infra
func (i *Infra) DescribeInfra(infra string) {
	client := newApiClient()

	response := client.action(path.Join(infraEntity, infra), "GET", nil)
	response.Process(true) // process response and exit if error

	// TODO: parse response.Body into required structure and return
}

// ListInfra : list all infra(s) with filters
func (i *Infra) ListInfra(team, state, env string) {
	client := newApiClient()
	client.QueryParams["team"] = team
	client.QueryParams["state"] = state
	client.QueryParams["env"] = env

	response := client.action(path.Join(infraEntity), "GET", nil)
	response.Process(true) // process response and exit if error

	// TODO: parse response.Body into required structure and return
}

// DeleteInfra : delete an created infra
func (i *Infra) DeleteInfra(infra string) {
	client := newApiClient()

	response := client.action(path.Join(infraEntity, infra), "DELETE", nil)
	response.Process(true) // process response and exit if error
	// no return required
}
