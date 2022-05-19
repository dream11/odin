package backend

import (
	"fmt"
	"github.com/dream11/odin/internal/config"
	"github.com/dream11/odin/internal/ui"
	"github.com/dream11/odin/pkg/request"
	"github.com/dream11/odin/pkg/sse"
	"time"
)

var logger ui.Logger

const (
	RetryCount     = 2
	BackoffSeconds = 1
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

func (c *clientProperties) actionWithRetry(entity, requestType string, body interface{}) request.Response {
	if RetryCount <= 0 {
		return c.action(entity, requestType, body)
	}

	var response request.Response

	backOffDuration := 0
	for i := 0; i < RetryCount; i++ {
		if backOffDuration > 0 {
			logger.Debug(fmt.Sprintf("Waiting for %d Second[s]", backOffDuration))
			time.Sleep(time.Duration(backOffDuration) * time.Second)
		}
		backOffDuration += BackoffSeconds
		response = c.action(entity, requestType, body)
		if response.Error == nil && response.StatusCode < 500 {
			return response
		}
	}
	return response
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

// initiation of an SSE client for backend stream interactions
type streamingClientProperties clientProperties

// perform streaming on initiated client
func (sc *streamingClientProperties) stream(entity, requestType string, body interface{}) sse.StreamResponse {
	req := sse.StreamRequest{
		Method: requestType,
		URL:    sc.address + entity,
		Header: sc.Headers,
		Query:  sc.QueryParams,
		Body:   body,
	}

	response := req.Stream()
	return response
}

func (sc *streamingClientProperties) streamWithRetry(entity, requestType string, body interface{}) sse.StreamResponse {
	if RetryCount <= 0 {
		return sc.stream(entity, requestType, body)
	}

	var response sse.StreamResponse

	backOffDuration := 0
	for i := 0; i <= RetryCount; i++ {
		if backOffDuration > 0 {
			logger.Debug(fmt.Sprintf("Waiting for %d Second[s]", backOffDuration))
			time.Sleep(time.Duration(backOffDuration) * time.Second)
		}
		backOffDuration += BackoffSeconds
		response = sc.stream(entity, requestType, body)
		if response.Error == nil && response.StatusCode < 500 {
			return response
		}
	}
	return response
}

func newStreamingClient() streamingClientProperties {
	var appConfig = config.Get()

	return streamingClientProperties{
		address: appConfig.BackendAddr + "/",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		QueryParams: map[string]string{},
	}
}

func newStreamingApiClient() streamingClientProperties {
	var appConfig = config.Get()

	streamClient := newStreamingClient()
	streamClient.address += "api/integration/cli/stream/v2/"
	streamClient.Headers["Authorization"] = "Bearer " + appConfig.AccessToken

	return streamClient
}
