package backend

import (
	"encoding/json"
	"path"

	"github.com/dream11/odin/api/component"
)

// Component entity
type Component struct{}

// DescribeComponent : describe a component type
func (c *Component) DescribeComponent(componentName, version string) (component.Component, error) {
	client := newApiClient()
	response := client.action(path.Join("components", componentName, "versions", version), "GET", nil)
	response.Process(true) // process response and exit if error

	var componentResponse component.DetailComponentResponse
	err := json.Unmarshal(response.Body, &componentResponse)

	return componentResponse.Response, err
}
