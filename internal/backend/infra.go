package backend

import (
	"encoding/json"
	"path"

	infraResp "github.com/dream11/odin/api/infra"
)

// Infra entity
type Infra struct{}

// root entity
var infraEntity = "infras"

// CreateInfra : create an empty infra
func (i *Infra) CreateInfra(infraDetails interface{}) (infraResp.CreationName, error) {
	client := newApiClient()

	response := client.action(infraEntity + "/", "POST", infraDetails)
	response.Process(true) // process response and exit if error

	var infraResponse infraResp.CreationResponse
	err := json.Unmarshal(response.Body, &infraResponse)

	return infraResponse.Response, err
}

// DescribeInfra : describe an infra
func (i *Infra) DescribeInfra(infra string) {
	client := newApiClient()

	response := client.action(path.Join(infraEntity, infra), "GET", nil)
	response.Process(true) // process response and exit if error

	// TODO: parse response.Body into required structure and return
}

// ListInfra : list all infra(s) with filters
func (i *Infra) ListInfra() ([]infraResp.Infra, error) {
	client := newApiClient()

	response := client.action(infraEntity + "/", "GET", nil)
	response.Process(true) // process response and exit if error

	var infraResponse infraResp.ListResponse
	err := json.Unmarshal(response.Body, &infraResponse)

	return infraResponse.Response, err
}

// DeleteInfra : delete an created infra
func (i *Infra) DeleteInfra(infra string) {
	client := newApiClient()

	response := client.action(path.Join(infraEntity, infra), "DELETE", nil)
	response.Process(true) // process response and exit if error
}

// UpdateInfra : update a created infra
func (i *Infra) UpdateInfra(infra string, config interface{}) {
	client := newApiClient()

	response := client.action(path.Join(infraEntity, infra) + "/", "PUT", config)
	response.Process(true) // process response and exit if error
}
