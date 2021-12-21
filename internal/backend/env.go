package backend

import (
	"encoding/json"
	"path"

	"github.com/dream11/odin/api/environment"
)

// Env entity
type Env struct{}

// root entity
var envEntity = "infras"

// CreateEnv : create an empty environment
func (e *Env) CreateEnv(envDetails interface{}) (environment.Env, error) {
	client := newApiClient()

	response := client.action(envEntity+"/", "POST", envDetails)
	response.Process(true) // process response and exit if error

	var envResponse environment.CreationResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// DescribeEnv : describe an environment
func (e *Env) DescribeEnv(env string) ([]environment.Env, error) {
	client := newApiClient()

	response := client.action(path.Join(envEntity, env)+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse environment.ListResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// ListEnv : list all environment(s) with filters
func (e *Env) ListEnv() ([]environment.Env, error) {
	client := newApiClient()

	response := client.action(envEntity+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse environment.ListResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// DeleteEnv : delete an created environment
func (e *Env) DeleteEnv(env string) {
	client := newApiClient()

	response := client.action(path.Join(envEntity, env)+"/", "DELETE", nil)
	response.Process(true) // process response and exit if error
}

// UpdateEnv : update a created environment
func (e *Env) UpdateEnv(env string, config interface{}) {
	client := newApiClient()

	response := client.action(path.Join(envEntity, env)+"/", "PUT", config)
	response.Process(true) // process response and exit if error
}
