package backend

import (
	"fmt"
	"path"
	"encoding/json"

	"github.com/dream11/odin/api/service"
)

// Service entity
type Service struct{}

// root entity
var serviceEntity = "services"

// CreateService : register a service version with backend
func (s *Service) CreateService(service interface{}) {
	client := newApiClient()

	response := client.action(serviceEntity + "/", "POST", service)
	response.Process(true) // process response and exit if error
}

// DescribeService : describe a service version or all versions of a service
func (s *Service) DescribeService(service, version string) {
	client := newApiClient()
	client.QueryParams["version"] = version

	response := client.action(path.Join(serviceEntity, service), "GET", nil)
	response.Process(true)

	// TODO: parse response.Body into required structure and return
}

// ListServices : list services per team and describe versions
func (s *Service) ListServices(team, version string, maturity bool) ([]service.Service, error) {
	client := newApiClient()
	client.QueryParams["team"] = team
	client.QueryParams["version"] = version
	client.QueryParams["isMature"] = fmt.Sprintf("%v", maturity)

	response := client.action(serviceEntity, "GET", nil)
	response.Process(true)

	var serviceResponse service.ListResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

// DeleteService : delete a service version
func (s *Service) DeleteService(service, version string) {
	client := newApiClient()

	response := client.action(path.Join(serviceEntity, service, "version", version) + "/", "DELETE", nil)
	response.Process(true)
}

// MarkMature : mark a service as mature
func (s *Service) MarkMature(service, version string) {
	client := newApiClient()

	response := client.action(path.Join(serviceEntity, service, "version", version, "mature"), "PUT", nil)
	response.Process(true)
}

// DeployService : deploy a service
func (s *Service) DeployService(service, version, infra, env string) {
	client := newApiClient()

	body := map[string]string{
		"infra_name": infra,
		"env":        env,
	}

	// TODO: validate no conversion required
	//jsonBody, err := json.Marshal(body)
	//if err != nil {
	//	logger.Error("Unable to format json body! " + err.Error())
	//	os.Exit(1)
	//}

	response := client.action(path.Join(serviceEntity, "deploy", service, "version", version), "POST", body)
	response.Process(true)
}
