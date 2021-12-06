package backend

import (
	"encoding/json"

	"github.com/dream11/odin/api/component"
)

// Component entity
type Component struct{}

// ListComponents : list all available component types
func (c *Component) ListComponents() ([]component.Type, error) {
	client := newApiClient()

	response := client.action("componenttypes", "GET", nil)
	response.Process(true) // process response and exit if error

	var componentTypeResponse component.ListTypeResponse
	err := json.Unmarshal(response.Body, &componentTypeResponse)

	return componentTypeResponse.Response, err
}
