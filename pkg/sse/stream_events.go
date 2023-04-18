package sse

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/apoorvam/goterminal"
	"github.com/briandowns/spinner"
	"github.com/dream11/odin/internal/ui"
	"github.com/dream11/odin/pkg/request"
	"github.com/r3labs/sse"
)

// StreamRequest structure
type StreamRequest request.Request

// StreamResponse structure
type StreamResponse request.Response

var logger ui.Logger

var SPINNER_COLOR = "fgHiBlue"
var SPINNER_STYLE = "bold"
var SPINNER_TYPE = 14
var SPINNER_DELAY = 100 * time.Millisecond

func (sr *StreamRequest) Stream() StreamResponse {
	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(sr.Body)
	if err != nil {
		logger.Debug(err.Error())
		return StreamResponse{Error: err}
	}

	req, err := http.NewRequest(sr.Method, sr.URL, payload)
	if err != nil {
		logger.Debug(err.Error())
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
		logger.Debug(err.Error())
		return StreamResponse{Error: err}
	}

	writer := goterminal.New(os.Stdout)
	ms := ui.New(spinner.CharSets[SPINNER_TYPE], SPINNER_DELAY, writer, SPINNER_TYPE)

	data := bufio.NewScanner(resp.Body)
	s := spinner.New(spinner.CharSets[SPINNER_TYPE], SPINNER_DELAY)
	for data.Scan() {
		line := string(data.Bytes())
		if line == "" {
			continue
		}
		if strings.Contains(line, ui.MULTISPINNER) {
			ms.Data = line
			s.Stop()
			ms.Start()
		} else if strings.Contains(line, ui.SPINNER) {
			parts := strings.Split(line, ui.SPINNER)
			s.Prefix = parts[0]
			s.Suffix = parts[1]
			s.HideCursor = false
			err := s.Color(SPINNER_COLOR, SPINNER_STYLE)
			if err != nil {
				logger.Error(err.Error())
			}
			ms.Stop()
			s.Start()
		} else {
			s.Stop()
			ms.Stop()
			logger.Output(line + "\n")
		}
		/**
		This is to evaluate if the streamed response from BE has any error information,
		if yes the idea is to extract the response code from the response data and use it to set status codes.
		Ex. Error message from BE is of the following format,
		ERROR(400): Cloud Provider Account 'adsdsasdassd' is not one of ['load', 'prod', 'proto', 'proto-dev', 'staging']
		*/
		if strings.HasPrefix(line, "ERROR") {
			re := regexp.MustCompile("\\d+")
			resp.StatusCode, _ = strconv.Atoi(re.FindAllString(line, 1)[0])
		}
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
			request.HandleExit(1, exitOnError)
		}
	}
}
