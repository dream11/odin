package backend

import (
	"encoding/json"
	"github.com/dream11/odin/api/component"
	"path"
)

// Component entity
type ComponentType struct{}

// ListComponentTypes : list all available component types
func (c *ComponentType) ListComponentTypes(componentTypeName, version string) ([]component.Type, error) {
	client := newApiClient()
	client.QueryParams["version"] = version
	client.QueryParams["name"] = componentTypeName
	response := client.action("componenttypes", "GET", nil)

	response.Process(true) // process response and exit if error

	var componentTypeResponse component.ListTypeResponse
	err := json.Unmarshal(response.Body, &componentTypeResponse)

	return componentTypeResponse.Response, err
}

// DescribeComponentTypes : describe a component type
func (c *ComponentType) DescribeComponentTypes(componentTypeName, version string) (component.Type, error) {
	client := newApiClient()
	client.QueryParams["version"] = version
	response := client.action(path.Join("componenttypes", componentTypeName), "GET", nil)
	response.Process(true) // process response and exit if error

	var componentTypeResponse component.DetailComponentTypeResponse
	err := json.Unmarshal(response.Body, &componentTypeResponse)

	return componentTypeResponse.Response, err
}
