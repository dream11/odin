package request

import (
	"bytes"
	"encoding/json"
	"github.com/dream11/odin/internal/ui"
	"io/ioutil"
	"net/http"
	"os"
)

// Request structure
type Request struct {
	Method string
	URL    string
	Header map[string]string
	Query  map[string]string
	Body   interface{}
}

// Response structure
type Response struct {
	Status     string
	StatusCode int
	Body       []byte
	Error      error
}

var logger ui.Logger

// Make : make a generated request
func (r *Request) Make() Response {
	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(r.Body)
	if err != nil {
		return Response{Error: err}
	}

	client := &http.Client{}
	request, err := http.NewRequest(r.Method, r.URL, payload)
	if err != nil {
		return Response{Error: err}
	}

	q := request.URL.Query()
	for key, val := range r.Query {
		if len(val) > 0 {
			q.Add(key, val)
		}
	}
	request.URL.RawQuery = q.Encode()

	logger.Debug("URL: " + request.URL.String())

	for key, value := range r.Header {
		if len(value) > 0 {
			request.Header.Set(key, value)
		}
	}

	// what if the internet is not present on the client side
	response, err := client.Do(request)

	if err != nil {
		return Response{Error: err}
	}

	respBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return Response{Error: err}
	}

	return Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Body:       respBody,
		Error:      nil,
	}
}

// Process : process request response to generate valid output
// Exit on error, only if specified
func (r *Response) Process(exitOnError bool) {
	// Parse error and display error message
	if r.Error != nil {
		logger.Error(r.Error.Error())
		handleExit(1, exitOnError)
	} else {
		if matchStatusCode(r.StatusCode, 200) {
			logger.Success(r.Status)
			logger.Debug(string(r.Body))
		} else if matchStatusCode(r.StatusCode, 300) {
			logger.Warn(r.Status)
			logger.Debug(string(r.Body))
		} else if matchStatusCode(r.StatusCode, 400) || matchStatusCode(r.StatusCode, 500) {
			logger.Error(r.Status)
			logger.Debug(string(r.Body))
			handleExit(1, exitOnError)
		}
	}
}

// handle exit calls, allow exit if allowed via boolean
func handleExit(code int, exit bool) {
	if exit {
		os.Exit(code)
	}
}

// match the status code range
func matchStatusCode(statusCode, matchCode int) bool {
	return (statusCode - matchCode) < 100
}
