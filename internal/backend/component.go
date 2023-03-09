package backend

import (
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
