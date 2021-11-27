package backend

// Component entity
type Component struct{}

// ListComponents : list all available component types
func (c *Component) ListComponents() {
	client := newApiClient()

	response := client.action("componenttypes", "GET", nil)
	response.Process(true) // process response and exit if error

	// TODO: parse response.Body into required structure and return
}
