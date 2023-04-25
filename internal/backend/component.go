package backend

import (
	"encoding/json"
	"path"

	"github.com/dream11/odin/api/component"
)

type Component struct{}

func (c *Component) OperateComponent(componentName string, data component.OperateComponentRequest) {
	client := newStreamingApiClient()
	client.Headers["Command-Verb"] = "operate"
	response := client.streamWithRetry(path.Join("component", componentName, "operate"), "PUT", data)
	response.Process(true)
}

func (c *Component) CompareOperationChanges(componentName string, data component.OperateComponentRequest) (interface{}, error){
	client := newApiClient()
	client.Headers["Command-Verb"] = "operate"
	response := client.actionWithRetry(path.Join("component", componentName, "compare"), "POST", data)
	response.Process(true)
	var compareResponse component.CompareOperationChangesResponse
	err := json.Unmarshal(response.Body, &compareResponse)

	return compareResponse.Response, err
}
