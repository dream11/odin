package backend

import (
	"encoding/json"
	"path"

	"github.com/dream11/odin/api/service"
)

// Service entity
type Service struct{}

// root entity
var serviceEntity = "services"

func (s *Service) CreateServiceStream(serviceDefinition interface{}, provisioningConfigMap map[string]interface{}) {
	client := newStreamingApiClient()
	response := client.stream(serviceEntity+"/", "POST", service.MergedService{Service: serviceDefinition, ProvisioningConfig: provisioningConfigMap})
	response.Process(true) // process response and exit if error
}

// RebuildServiceStream : rebuild a service using streams
func (s *Service) RebuildServiceStream(service, version string) {
	client := newStreamingApiClient()

	response := client.stream(path.Join(serviceEntity, service, "versions", version, "rebuild")+"/", "PUT", nil)
	response.Process(true)
}

// DescribeService : describe a service version or all versions of a service
func (s *Service) DescribeService(name, version, component string) (service.Service, error) {
	client := newApiClient()
	client.QueryParams["version"] = version
	client.QueryParams["component"] = component
	response := client.actionWithRetry(path.Join(serviceEntity, name), "GET", nil)
	response.Process(true)

	var serviceResponse service.DetailResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

// ListServices : list services per team and describe versions
func (s *Service) ListServices(team, version, serviceName string, label string) ([]service.Service, error) {
	client := newApiClient()
	client.QueryParams["team"] = team
	client.QueryParams["version"] = version
	client.QueryParams["name"] = serviceName
	client.QueryParams["label"] = label

	response := client.actionWithRetry(serviceEntity, "GET", nil)
	response.Process(true)

	var serviceResponse service.ListResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

// UnDeployServiceStream : un-deploy a service in an Env and stream creation events
func (s *Service) UnDeployServiceStream(serviceName, envName string) {
	client := newStreamingApiClient()
	client.QueryParams["env_name"] = envName

	response := client.streamWithRetry(path.Join(serviceEntity, "undeploy", serviceName)+"/", "DELETE", nil)
	response.Process(true)
}

// DeleteService : delete a service version
func (s *Service) DeleteService(service, version string) {
	client := newApiClient()

	response := client.actionWithRetry(path.Join(serviceEntity, service, "versions", version)+"/", "DELETE", nil)
	response.Process(true)
}

// LabelService : label a service
func (s *Service) LabelService(service, version, label string) {
	client := newApiClient()

	data := map[string]interface{}{
		"resource-name":    service,
		"resource-version": version,
		"label":            label,
	}
	response := client.actionWithRetry(path.Join(serviceEntity, service, "version", version, "label")+"/", "PUT", data)
	response.Process(true)
}

// UnlabelService : unlabel a service
func (s *Service) UnlabelService(service, version, label string) {
	client := newApiClient()

	data := map[string]interface{}{
		"resource-name":    service,
		"resource-version": version,
		"label":            label,
	}
	response := client.actionWithRetry(path.Join(serviceEntity, service, "version", version, "unlabel")+"/", "PUT", data)
	response.Process(true)
}

// DeployServiceStream : deploy a service in an Env and stream creation events
func (s *Service) DeployReleasedServiceStream(service, version, env, configStoreNamespace string, provisionConfig interface{}) {
	client := newStreamingApiClient()
	client.QueryParams["env_name"] = env
	client.QueryParams["config_store_namespace"] = configStoreNamespace

	data := map[string]interface{}{}

	if provisionConfig != nil {
		data["provisionConfig"] = provisionConfig
	}

	response := client.streamWithRetry(path.Join(serviceEntity, "deploy", service, "versions", version)+"/", "POST", data)
	response.Process(true)
}

// DeployServiceStream : deploy a service in an Env and stream creation events
func (s *Service) DeployUnreleasedServiceStream(serviceDefinition, provisionConfig interface{}, env, configStoreNamespace string) {
	client := newStreamingApiClient()
	client.QueryParams["env_name"] = env
	client.QueryParams["config_store_namespace"] = configStoreNamespace

	data := map[string]interface{}{
		"serviceDefinition": serviceDefinition,
	}

	if provisionConfig != nil {
		data["provisionConfig"] = provisionConfig
	}

	response := client.streamWithRetry(path.Join(serviceEntity, "deploy")+"/", "POST", data)
	response.Process(true)
}

// CompareReleasedService : Find diff in service
func (s *Service) CompareService(envName, service, version, serviceDefinition interface{}, parsedProvisioningConfig interface{}, configStoreNamespace *string) {
	client := newApiClient()
	data := map[string]interface{}{
		"serviceDefinition":  serviceDefinition,
		"provisioningConfig":  parsedProvisioningConfig,
		"configStoreNamespace":  configStoreNamespace,
		"serviceName":        service,
		"serviceVersion":        version,
		"envName":        envName,
	}
	response := client.actionWithRetry(path.Join(serviceEntity, "compare"), "GET", data)
	response.Process(true)
}

// StatusService : get status of a service
func (s *Service) StatusService(serviceName, version string) ([]service.Status, error) {
	client := newApiClient()

	response := client.actionWithRetry(path.Join(serviceEntity, serviceName, "versions", version, "status")+"/", "GET", nil)
	response.Process(true)

	var serviceResponse service.StatusResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}
