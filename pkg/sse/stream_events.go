package sse

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dream11/odin/internal/ui"
	"github.com/dream11/odin/pkg/request"
	"github.com/r3labs/sse"
)

// StreamRequest structure
type StreamRequest request.Request

// StreamResponse structure
type StreamResponse request.Response

var logger ui.Logger

func (sr *StreamRequest) Stream() StreamResponse {
	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(sr.Body)
	if err != nil {
		logger.Error(err.Error())
		return StreamResponse{Error: err}
	}

	req, err := http.NewRequest(sr.Method, sr.URL, payload)
	if err != nil {
		logger.Error(err.Error())
		return StreamResponse{Error: err}
	}

	q := req.URL.Query()
	for key, val := range sr.Query {
		if len(val) > 0 {
			q.Add(key, val)
		}
	}

	req.URL.RawQuery = q.Encode()

	logger.Debug("Stream URL: " + req.URL.String())

	for key, value := range sr.Header {
		if len(value) > 0 {
			req.Header.Set(key, value)
		}
	}

	sseClient := sse.NewClient(sr.URL)

	resp, err := sseClient.Connection.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return StreamResponse{Error: err}
	}

	data := bufio.NewScanner(resp.Body)
	for data.Scan() {
		line := data.Bytes()
		logger.Output(string(line))
	}

	return StreamResponse{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Body:       nil,
		Error:      nil,
	}
}

// Process : process request response to generate valid output
// Exit on error, only if specified
func (r *StreamResponse) Process(exitOnError bool) {
	// Parse error and display error message
	if r.Error != nil {
		logger.Error(r.Error.Error())
		request.HandleExit(1, exitOnError)
	} else {
		if request.MatchStatusCode(r.StatusCode, 200) {
			logger.Debug(r.Status)
		} else if request.MatchStatusCode(r.StatusCode, 300) {
			logger.Debug(fmt.Sprintf("[%d] %s", r.StatusCode, r.Status))
			logger.Debug(string(r.Body))
		} else if request.MatchStatusCode(r.StatusCode, 400) || request.MatchStatusCode(r.StatusCode, 500) {
			logger.Debug(r.Status)
			logger.Error(string(r.Body))
			request.HandleExit(1, exitOnError)
		}
	}
}
