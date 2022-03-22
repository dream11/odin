package backend

import (
	"encoding/json"
	"github.com/dream11/odin/api/servicegroup"
)

// Service entity
type ServiceGroup struct{}

// root entity
var serviceGroupEntity = "servicegroup"

// CreateServiceGroup : register a service-group with backend
func (s *ServiceGroup) CreateServiceGroup(serviceGroup interface{}) (string, error) {
	client := newApiClient()

	response := client.action(serviceGroupEntity+"/", "POST", serviceGroup)
	response.Process(true) // process response and exit if error

	var serviceResponse servicegroup.CreateResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response.Message, err
}
