package backend

import (
	"encoding/json"
	"github.com/dream11/odin/api/servicegroup"
	"path"
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

// ListServices : list services per team and describe versions
func (s *ServiceGroup) ListServiceGroups(serviceGroupName, serviceName string) ([]servicegroup.List, error) {
	client := newApiClient()
	client.QueryParams["name"] = serviceGroupName
	client.QueryParams["service"] = serviceName

	response := client.action(serviceGroupEntity, "GET", nil)
	response.Process(true)

	var serviceResponse servicegroup.ListResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *ServiceGroup) DescribeService(serviceGroupName string) (servicegroup.Describe, error) {
	client := newApiClient()

	response := client.action(path.Join(serviceGroupEntity, serviceGroupName), "GET", nil)
	response.Process(true)

	var serviceResponse servicegroup.DescribeResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}
