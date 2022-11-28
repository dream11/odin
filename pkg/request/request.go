package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dream11/odin/internal/ui"
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

	respBody, err := io.ReadAll(response.Body)

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
		HandleExit(1, exitOnError)
	} else {
		if MatchStatusCode(r.StatusCode, 200) {
			logger.Debug(r.Status)
			logger.Debug(string(r.Body))
		} else if MatchStatusCode(r.StatusCode, 300) {
			logger.Debug(r.Status)
			logger.Debug(string(r.Body))
		} else if MatchStatusCode(r.StatusCode, 400) || MatchStatusCode(r.StatusCode, 500) {
			logger.Debug(r.Status)
			logger.Error(string(r.Body))
			HandleExit(1, exitOnError)
		}
	}
}

// Process Request Response and Return Status Code

func (r *Response) GetStatusCode(exitOnError bool) int {

	if r.Error != nil {
		logger.Error(r.Error.Error())
		HandleExit(1, exitOnError)
	} else {
		if MatchStatusCode(r.StatusCode, 200) {
			logger.Debug(r.Status)
			logger.Debug(string(r.Body))
			return 200
		} else if MatchStatusCode(r.StatusCode, 300) {
			logger.Debug(r.Status)
			logger.Debug(string(r.Body))
			return 300
		} else if MatchStatusCode(r.StatusCode, 400) {
			logger.Debug(r.Status)
			logger.Debug(string(r.Body))
			return 400
		} else if MatchStatusCode(r.StatusCode, 500) {
			logger.Debug(r.Status)
			logger.Debug(string(r.Body))
			return 500
		}
	}
	return 0
}

// Process : process request response to generate valid output
// Exit on error, only if specified
func (r *Response) ProcessHandleError(exitOnError bool) {
	// Parse error and display error message
	if r.Error != nil {
		logger.Error(r.Error.Error())
		HandleExit(1, exitOnError)
	} else {
		if MatchStatusCode(r.StatusCode, 200) {
			logger.Debug(r.Status)
		} else if MatchStatusCode(r.StatusCode, 300) {
			logger.Debug(r.Status)
			logger.Debug(string(r.Body))
		} else if MatchStatusCode(r.StatusCode, 400) || MatchStatusCode(r.StatusCode, 500) {
			logger.Debug(r.Status)
			handleError(r)
			HandleExit(1, exitOnError)
		}
	}
}

// HandleExit : handle exit calls, allow exit if allowed via boolean
func HandleExit(code int, exit bool) {
	if exit {
		os.Exit(code)
	}
}

// MatchStatusCode : match the status code range
func MatchStatusCode(statusCode, matchCode int) bool {
	return (statusCode - matchCode) < 100
}

func handleError(r *Response) {
	data := map[string]interface{}{}
	if err := json.Unmarshal(r.Body, &data); err != nil {
		logger.Error("Error while parsing error message: " + string(r.Body))
		return
	}
	if val, ok := data["error"]; ok {
		logger.Error(fmt.Sprintf("%v", val))
		return
	}
	logger.Error(string(r.Body))
}
