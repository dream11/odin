package backend

import (
	"encoding/json"
	"path"

	operationapi "github.com/dream11/odin/api/operation"
)

type Operation struct{}

func (o *Operation) ListOperations(componentTypeName string) ([]operationapi.Operation, error) {
	client := newApiClient()
	response := client.actionWithRetry(path.Join("component", componentTypeName, "operate"), "GET", nil)
	response.Process(true) // process response and exit if error
	var listResponse operationapi.ListOperation
	err := json.Unmarshal(response.Body, &listResponse)
	return listResponse.Response, err
}
