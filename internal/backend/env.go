package backend

import (
	"encoding/json"
	"path"

	envResp "github.com/dream11/odin/api/environment"
)

// Env entity
type Env struct{}

// root entity
var envEntity = "envs"

// CreateEnv : create an empty Env
func (e *Env) CreateEnv(envDetails interface{}) (envResp.Env, error) {
	client := newApiClient()

	response := client.action(envEntity+"/", "POST", envDetails)
	response.Process(true) // process response and exit if error

	var envResponse envResp.CreationResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// DescribeEnv : describe an Env
func (e *Env) DescribeEnv(env, service, component string) (envResp.Env, error) {
	client := newApiClient()
	client.QueryParams["service"] = service
	client.QueryParams["component"] = component
	response := client.action(path.Join(envEntity, env)+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.DetailResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// ListEnv : list all environment(s) with filters
func (e *Env) ListEnv(name, team, env, providerAccount string) ([]envResp.Env, error) {
	client := newApiClient()
	client.QueryParams["name"] = name
	client.QueryParams["team"] = team
	client.QueryParams["envType"] = env
	client.QueryParams["cloudProviderAccount"] = providerAccount
	response := client.action(envEntity+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.ListResponse
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

// GetHistoryEnv : get historical changes in an Env
func (e *Env) GetHistoryEnv(env string) ([]envResp.History, error) {
	client := newApiClient()

	response := client.action(path.Join("envhistory", env)+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.HistoryListResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// DescribeHistoryEnv : describe a historical changes in an Env
func (e *Env) DescribeHistoryEnv(env string, id string) ([]envResp.History, error) {
	client := newApiClient()
	client.QueryParams["id"] = id

	response := client.action(path.Join("envhistory", env)+"/", "GET", nil)
	response.Process(false) // process response and exit if error

	var envResponse envResp.HistoryListResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}
