package backend

import (
	"encoding/json"

	envTypeResp "github.com/dream11/odin/api/envtype"
)

// Env entity
type EnvType struct{}

// root entity
var envEntityType = "envtypes"

// ListEnv : list all environment(s) with filters
func (e *Env) ListEnvType() ([]string, error) {
	client := newApiClient()
	response := client.actionWithRetry(envEntityType+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envTypeResponse envTypeResp.ListResponse
	err := json.Unmarshal(response.Body, &envTypeResponse)

	return envTypeResponse.Response, err
}
