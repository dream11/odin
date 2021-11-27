package backend

import (
	"github.com/dream11/odin/internal/config"
	"github.com/dream11/odin/internal/ui"
	"github.com/dream11/odin/pkg/request"
)

type clientProperties struct {
	address     string
	Headers     map[string]string
	QueryParams map[string]string
}

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

var appConfig = config.Get()
var logger ui.Logger

func newClient() clientProperties {
	return clientProperties{
		address: appConfig.BackendAddr + "/",
		Headers: map[string]string{
			"Content-Type":   "text/JSON",
		},
		QueryParams: map[string]string{},
	}
}

func newApiClient() clientProperties {
	apiClient := newClient()
	apiClient.address += "api/integration/cli/v1/"
	apiClient.Headers["Authentication"] = "Bearer " + appConfig.AccessToken

	return apiClient
}
