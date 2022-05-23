package backend

import (
	"encoding/json"
	"path"

	"github.com/dream11/odin/api/component"
)

// ComponentType entity
type ComponentType struct{}

// ListComponentTypes : list all available component types
func (c *ComponentType) ListComponentTypes(componentTypeName, version string) ([]component.Type, error) {
	client := newApiClient()
	client.QueryParams["version"] = version
	client.QueryParams["name"] = componentTypeName
	response := client.actionWithRetry("componenttypes", "GET", nil)

	response.Process(true) // process response and exit if error

	var componentTypeResponse component.ListTypeResponse
	err := json.Unmarshal(response.Body, &componentTypeResponse)

	return componentTypeResponse.Response, err
}

// DescribeComponentTypes : describe a component type
func (c *ComponentType) DescribeComponentType(componentTypeName, version string) (component.ComponentDetails, error) {
	client := newApiClient()
	client.QueryParams["version"] = version
	response := client.actionWithRetry(path.Join("componenttypes", componentTypeName), "GET", nil)
	response.Process(true) // process response and exit if error

	var componentDetailsResponse component.ComponentDetailsResponse
	err := json.Unmarshal(response.Body, &componentDetailsResponse)

	return componentDetailsResponse.Response, err
}
