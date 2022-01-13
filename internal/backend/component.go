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
func (c *ComponentType) DescribeComponentType(componentTypeName, version string) (component.Type, error) {
	client := newApiClient()
	client.QueryParams["version"] = version
	response := client.action(path.Join("componenttypes", componentTypeName), "GET", nil)
	response.Process(true) // process response and exit if error

	var componentTypeResponse component.DetailComponentTypeResponse
	err := json.Unmarshal(response.Body, &componentTypeResponse)

	return componentTypeResponse.Response, err
}


// DescribeComponent : describe a component type
func (c *ComponentType) DescribeComponent(componentName, version string) (component.Component, error) {
	client := newApiClient()
	client.QueryParams["version"] = version
	response := client.action(path.Join("components", componentName, "versions", version), "GET", nil)
	response.Process(true) // process response and exit if error

	var componentResponse component.DetailComponentResponse
	err := json.Unmarshal(response.Body, &componentResponse)

	return componentResponse.Response, err
}