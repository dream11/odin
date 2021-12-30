package backend

import (
	"encoding/json"
	"path"

	envResp "github.com/dream11/odin/api/env"
)

// Env entity
type Env struct{}

// root entity
var envEntity = "Envs"

// CreateEnv : create an empty Env
func (i *Env) CreateEnv(EnvDetails interface{}) (envResp.Env, error) {
	client := newApiClient()

	response := client.action(envEntity+"/", "POST", EnvDetails)
	response.Process(true) // process response and exit if error

	var envResponse envResp.CreationResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// DescribeEnv : describe an Env
func (i *Env) DescribeEnv(Env string) ([]envResp.Env, error) {
	client := newApiClient()

	response := client.action(path.Join(envEntity, Env)+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.ListResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// ListEnv : list all Env(s) with filters
func (i *Env) ListEnv() ([]envResp.Env, error) {
	client := newApiClient()

	response := client.action(envEntity+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.ListResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// DeleteEnv : delete an created Env
func (i *Env) DeleteEnv(Env string) {
	client := newApiClient()

	response := client.action(path.Join(envEntity, Env)+"/", "DELETE", nil)
	response.Process(true) // process response and exit if error
}

// UpdateEnv : update a created Env
func (i *Env) UpdateEnv(Env string, config interface{}) {
	client := newApiClient()

	response := client.action(path.Join(envEntity, Env)+"/", "PUT", config)
	response.Process(true) // process response and exit if error
}
