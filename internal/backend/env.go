package backend

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	envResp "github.com/dream11/odin/api/environment"
	"github.com/dream11/odin/internal/config"
	"github.com/dream11/odin/pkg/file"
)

// Env entity
type Env struct{}

// root entity
var envEntity = "envs"

// CreateEnv : create an empty Env
func (e *Env) CreateEnv(envDetails interface{}) (envResp.Env, error) {
	client := newApiClient()

	response := client.actionWithRetry(envEntity+"/", "POST", envDetails)
	response.Process(true) // process response and exit if error

	var envResponse envResp.CreationResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// CreateEnvStream : create an empty Env and stream creation events
func (e *Env) CreateEnvStream(envDetails interface{}) {
	client := newStreamingApiClient()
	response := client.streamWithRetry(envEntity+"/", "POST", envDetails)
	response.Process(true)
}

// DescribeEnv : describe an Env
func (e *Env) DescribeEnv(env, service, component string) (envResp.Env, error) {
	client := newApiClient()
	client.QueryParams["service"] = service
	client.QueryParams["component"] = component
	response := client.actionWithRetry(path.Join(envEntity, env)+"/", "GET", nil)
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
	response := client.actionWithRetry(envEntity+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.ListResponse
	err := json.Unmarshal(response.Body, &envResponse)
	return envResponse.Response, err
}

// DeleteEnv : delete an created environment
func (e *Env) DeleteEnv(env string) (envResp.EnvDeleteResponse, error) {
	client := newApiClient()

	response := client.actionWithRetry(path.Join(envEntity, env)+"/", "DELETE", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.EnvDeleteResponse
	err := json.Unmarshal(response.Body, &envResponse)
	return envResponse, err
}

// DeleteEnvStream : delete a created environment and stream deletion events
func (e *Env) DeleteEnvStream(env string) {
	client := newStreamingApiClient()
	response := client.streamWithRetry(path.Join(envEntity, env)+"/", "DELETE", nil)
	response.Process(true)
}

// UpdateEnv : update a created environment
func (e *Env) UpdateEnv(env string, config interface{}) {
	client := newApiClient()

	response := client.actionWithRetry(path.Join(envEntity, env)+"/", "PUT", config)
	response.Process(true) // process response and exit if error
}

// GetHistoryEnv : get historical changes in an Env
func (e *Env) GetHistoryEnv(env string) ([]envResp.History, error) {
	client := newApiClient()

	response := client.actionWithRetry(path.Join("envhistory", env)+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.HistoryListResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// DescribeHistoryEnv : describe a historical changes in an Env
func (e *Env) DescribeHistoryEnv(env string, id string) ([]envResp.History, error) {
	client := newApiClient()
	client.QueryParams["id"] = id

	response := client.actionWithRetry(path.Join("envhistory", env)+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.HistoryListResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.Response, err
}

// EnvStatus : Fetch status of the env
func (e *Env) EnvStatus(env string) (envResp.EnvStatus, error) {
	client := newApiClient()

	response := client.actionWithRetry(path.Join(envEntity, env)+"/status", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.EnvStatusResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.EnvResponse, err
}

func (e *Env) EnvServiceStatus(env, serviceName string) (envResp.EnvServiceStatus, error) {
	client := newApiClient()

	response := client.actionWithRetry(path.Join(envEntity, env)+"/services/"+serviceName+"/status", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.EnvServiceStatusResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse.ServiceResponse, err
}

func (e *Env) EnvTypes() (envResp.EnvTypesResponse, error) {
	client := newApiClient()

	response := client.actionWithRetry("envtypes", "GET", nil)
	response.Process(true) // process response and exit if error

	var envResponse envResp.EnvTypesResponse
	err := json.Unmarshal(response.Body, &envResponse)

	return envResponse, err
}

func (e *Env) SetEnv(envName string) error {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	data, err := file.Read(dirname + "/.odin/config")
	if err != nil {
		return err
	}
	result := ""
	r, _ := regexp.Compile(`(?:odinEnvName: [a-zA-Z]+-\w+)`)
	match := r.FindString(string(data))
	if match != "" {
		result = strings.Replace(string(data), match, fmt.Sprintf("odinEnvName: %s", envName), 1)
	} else {
		result = string(data) + fmt.Sprintf("odinEnvName: %s\n", envName)
	}
	err = file.Write(dirname+"/.odin/config", result, 0755)
	if err != nil {
		return err
	}
	return err
}

func (e *Env) FetchSetEnv(envName string) string {
	if envName != "" {
		return envName
	} else {
		var appConfig = config.Get()
		return appConfig.OdinEnvName
	}
}
