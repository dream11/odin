package backend

import (
	"encoding/json"
	"os"
	"path"
)

// Service entity
type Service struct{}

// root entity
var serviceEntity = "services"

// CreateService : register a service version with backend
func (s *Service) CreateService(service []byte) {
	client := newApiClient()

	response := client.action(serviceEntity, "POST", service)
	response.Process(true) // process response and exit if error
	// no return required
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
func (s *Service) ListServices(team, version string) {
	client := newApiClient()
	client.QueryParams["team"] = team
	client.QueryParams["version"] = version

	response := client.action(serviceEntity, "GET", nil)
	response.Process(true)

	// TODO: parse response.Body into required structure and return
}

// DeleteService : delete a service version
func (s *Service) DeleteService(service, version string) {
	client := newApiClient()

	response := client.action(path.Join(serviceEntity, service, "version", version), "DELETE", nil)
	response.Process(true)
	// no return required
}

// MarkMature : mark a service as mature
func (s *Service) MarkMature(service, version string) {
	client := newApiClient()

	response := client.action(path.Join(serviceEntity, service, "version", version, "mature"), "PUT", nil)
	response.Process(true)
	// no return required
}

// DeployService : deploy a service
func (s *Service) DeployService(service, version, infra, env string) {
	client := newApiClient()

	body := map[string]string{
		"infra_name":     infra,
		"runtime_config": env,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		logger.Error("Unable to format json body! " + err.Error())
		os.Exit(1)
	}

	response := client.action(path.Join(serviceEntity, "deploy", service, "version", version), "POST", jsonBody)
	response.Process(true)
	// no return required
}