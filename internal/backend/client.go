package backend

import (
	"github.com/dream11/odin/internal/config"
	"github.com/dream11/odin/pkg/request"
)

// initiation of an HTTP client for backend interactions
type clientProperties struct {
	address     string
	Headers     map[string]string
	QueryParams map[string]string
}

// perform HTTP actions on initiated client
func (c *clientProperties) action(entity, requestType string, body interface{}) request.Response {
	// TODO: add auth token to required header key
	req := request.Request{
		Method: requestType,
		URL:    c.address + entity,
		Query:  c.QueryParams,
		Header: c.Headers,
		Body:   body,
	}

	return req.Make()
}

// initiate a functional backend base-client
func newClient() clientProperties {
	var appConfig = config.Get()

	return clientProperties{
		address: appConfig.BackendAddr + "/",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		QueryParams: map[string]string{},
	}
}

// initiate an API integration client on top of base-client
func newApiClient() clientProperties {
	var appConfig = config.Get()

	apiClient := newClient()
	apiClient.address += "api/integration/cli/v2/"
	apiClient.Headers["Authorization"] = "Bearer " + appConfig.AccessToken

	return apiClient
}
