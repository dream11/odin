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

// UndeployService: To remove a service from a given env
func (s *Service) UndeployService(serviceName, env_name string) {
	client := newApiClient()
	client.QueryParams["env_name"] = env_name

	response := client.action(path.Join(serviceEntity, "undeploy", serviceName)+"/", "DELETE", nil)
	response.Process(true)

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

	response := client.action(path.Join(serviceEntity, service, "versions", version, "mature")+"/", "PUT", nil)
	response.Process(true)
}

// DeployService : deploy a service
func (s *Service) DeployService(service, version, env string, force, rebuild bool) {
	client := newApiClient()
	client.QueryParams["env_name"] = env
	client.QueryParams["force"] = fmt.Sprintf("%v", force)
	client.QueryParams["rebuild"] = fmt.Sprintf("%v", rebuild)

	response := client.action(path.Join(serviceEntity, "deploy", service, "versions", version)+"/", "POST", nil)
	response.Process(true)
}

// StatusService : get status of a service
func (s *Service) StatusService(serviceName, version string) ([]service.Status, error) {
	client := newApiClient()

	response := client.action(path.Join(serviceEntity, serviceName, "versions", version, "status")+"/", "GET", nil)
	response.Process(true)

	var serviceResponse service.StatusResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}
