package backend

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/dream11/odin/api/service"
)

// Service entity
type Service struct{}

// root entity
var serviceEntity = "services"

// CreateService : register a service version with backend
func (s *Service) CreateService(service interface{}) {
	client := newApiClient()

	response := client.action(serviceEntity+"/", "POST", service)
	response.Process(true) // process response and exit if error
}

// Rebuild Service : rebuild a service
func (s *Service) RebuildService(service, version string) {
	client := newApiClient()

	response := client.action(path.Join(serviceEntity, service, "versions", version, "rebuild")+"/", "PUT", nil)
	response.Process(true)
}

// DescribeService : describe a service version or all versions of a service
func (s *Service) DescribeService(name, version string) (service.Service, error) {
	client := newApiClient()
	client.QueryParams["version"] = version

	response := client.action(path.Join(serviceEntity, name), "GET", nil)
	response.Process(true)

	var serviceResponse service.DetailResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

// ListServices : list services per team and describe versions
func (s *Service) ListServices(team, version, serviceName string, maturity bool) ([]service.Service, error) {
	client := newApiClient()
	client.QueryParams["team"] = team
	client.QueryParams["version"] = version
	client.QueryParams["name"] = serviceName
	// if maturity then only pass isMature in query params
	if maturity {
		client.QueryParams["isMature"] = fmt.Sprintf("%v", maturity)
	}

	response := client.action(serviceEntity, "GET", nil)
	response.Process(true)

	var serviceResponse service.ListResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

// DeleteService : delete a service version
func (s *Service) DeleteService(service, version string) {
	client := newApiClient()

	response := client.action(path.Join(serviceEntity, service, "versions", version)+"/", "DELETE", nil)
	response.Process(true)
}

// MarkMature : mark a service as mature
func (s *Service) MarkMature(service, version string) {
	client := newApiClient()

	response := client.action(path.Join(serviceEntity, service, "version", version, "mature")+"/", "PUT", nil)
	response.Process(true)
}

// DeployService : deploy a service
func (s *Service) DeployService(service, version, env string, config interface{}) {
	client := newApiClient()
	client.QueryParams["infra_name"] = env

	response := client.action(path.Join(serviceEntity, "deploy", service, "version", version)+"/", "POST", config)
	response.Process(true)
}
