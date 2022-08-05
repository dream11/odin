package backend

import (
	"encoding/json"
	"github.com/dream11/odin/api/envtype"
)

// Env entity
type EnvType struct{}

// root entity
var envEntityType = "envtypes"

// ListEnv : list all environment(s) with filters
func (e *EnvType) ListEnvType() ([]envtype.Type, error) {
	client := newApiClient()
	response := client.actionWithRetry(envEntityType+"/", "GET", nil)
	response.Process(true) // process response and exit if error

	var envTypeResponse envtype.ListTypeResponse
	err := json.Unmarshal(response.Body, &envTypeResponse)

	return envTypeResponse.Response, err
}
