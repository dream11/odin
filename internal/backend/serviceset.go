package backend

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/dream11/odin/api/serviceset"
)

// ServiceSet entity
type ServiceSet struct{}

// root entity
var serviceSetEntity = "serviceset"

func (s *ServiceSet) CreateServiceSet(serviceSetDefinition interface{}) {
	client := newApiClient()

	response := client.action(serviceSetEntity+"/", "POST", serviceSetDefinition)
	response.ProcessHandleError(true)
}

func (s *ServiceSet) ListServiceSet(serviceSetName, serviceName string) ([]serviceset.List, error) {
	client := newApiClient()
	client.QueryParams["name"] = serviceSetName
	client.QueryParams["service"] = serviceName

	response := client.action(serviceSetEntity, "GET", nil)
	response.Process(true)

	var serviceResponse serviceset.ListResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *ServiceSet) DescribeServiceSet(serviceSetName string) (serviceset.Describe, error) {
	client := newApiClient()

	response := client.action(path.Join(serviceSetEntity, serviceSetName), "GET", nil)
	response.Process(true)

	var serviceResponse serviceset.DescribeResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *ServiceSet) DeleteServiceSet(serviceSetName string) {
	client := newApiClient()

	response := client.action(path.Join(serviceSetEntity, serviceSetName)+"/", "DELETE", nil)
	response.Process(true)
}

func (s *ServiceSet) DeployServiceSet(serviceSetName, env, platform, configStoreNamespace string, forceDeployServices []serviceset.ListEnvService, force bool) ([]serviceset.ServiceSetDeploy, error) {
	client := newApiClient()
	client.QueryParams["env_name"] = env
	client.QueryParams["platform"] = platform
	client.QueryParams["force"] = fmt.Sprintf("%v", force)
	client.QueryParams["config_store_namespace"] = configStoreNamespace

	data := map[string]interface{}{
		"forceDeployServices": forceDeployServices,
	}

	response := client.action(path.Join(serviceSetEntity, "deploy", serviceSetName, "env", env)+"/", "POST", data)
	response.Process(true)

	var serviceResponse serviceset.ServiceSetDeployResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s ServiceSet) ListEnvServices(serviceSetName, env, filterBy string) ([]serviceset.ListEnvService, error) {
	client := newApiClient()
	client.QueryParams["filter_by"] = filterBy

	response := client.action(path.Join(serviceSetEntity, serviceSetName, "env", env, "service")+"/", "GET", nil)
	response.Process(true)

	var serviceResponse serviceset.ListEnvServiceResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *ServiceSet) UndeployServiceSet(serviceSetName, env string, forceUndeployServices []serviceset.ListEnvService, force bool) ([]serviceset.ServiceSetDeploy, error) {
	client := newApiClient()
	client.QueryParams["env_name"] = env
	client.QueryParams["force"] = fmt.Sprintf("%v", force)

	data := map[string]interface{}{
		"forceUndeployServices": forceUndeployServices,
	}

	response := client.action(path.Join(serviceSetEntity, "undeploy", serviceSetName, "env", env)+"/", "DELETE", data)
	response.Process(true)

	var serviceResponse serviceset.ServiceSetDeployResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *ServiceSet) UpdateServiceSet(serviceSetName string, serviceSet interface{}) {
	client := newApiClient()

	response := client.action(path.Join(serviceSetEntity, serviceSetName)+"/", "PUT", serviceSet)
	response.Process(true) // process response and exit if error
}
