package backend

import (
	"encoding/json"
	labelapi "github.com/dream11/odin/api/label"
)

type Label struct{}
var labelRoot = "labels"

func (l *Label) CreateLabel(labelDetails interface{}) {
	client := newApiClient()
	response := client.action(labelRoot+"/", "POST", labelDetails)
	response.Process(true) // process response and exit if error
}

func (l *Label) ListLables() ([]labelapi.Label, error){
	client := newApiClient()
	response := client.action(labelRoot+"/", "GET", nil)
	response.Process(true) // process response and exit if error
	var listResponse labelapi.ListResponse
	err := json.Unmarshal(response.Body, &listResponse)
	return listResponse.Response, err
	
}

func (l *Label) DeleteLabel(label string) {
	client := newApiClient()
	client.QueryParams["name"] = label
	response := client.action(labelRoot+"/", "DELETE", nil)
	response.Process(true) // process response and exit if error
}

