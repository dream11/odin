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

func (s *ServiceSet) CreateServiceSet(serviceSetDefinition interface{}) (string, error) {
	client := newApiClient()

	response := client.action(serviceSetEntity+"/", "POST", serviceSetDefinition)
	response.Process(true) // process response and exit if error

	var serviceResponse serviceset.CreateResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response.Message, err
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

func (s *ServiceSet) DeployServiceSet(serviceSetName, env, platform string, forceDeployServices []serviceset.ListEnvService, force bool) ([]serviceset.ServiceSetServiceDeploy, error) {
	client := newApiClient()
	client.QueryParams["env_name"] = env
	client.QueryParams["platform"] = platform
	client.QueryParams["force"] = fmt.Sprintf("%v", force)

	response := client.action(path.Join(serviceSetEntity, "deploy", serviceSetName, "env", env)+"/", "POST", forceDeployServices)
	response.Process(true)

	var serviceResponse serviceset.ServiceSetServiceDeployResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s ServiceSet) ListEnvServices(serviceSetName, env string, isConflicted bool) ([]serviceset.ListEnvService, error) {
	client := newApiClient()
	client.QueryParams["isConflicted"] = fmt.Sprintf("%v", isConflicted)

	response := client.action(path.Join(serviceSetEntity, serviceSetName, "env", env, "service")+"/", "GET", nil)
	response.Process(true)

	var serviceResponse serviceset.ListEnvServiceResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *ServiceSet) UndeployServiceSet(serviceSetName, env string, forceDeployServices []serviceset.ListEnvService, force bool) ([]serviceset.ServiceSetServiceDeploy, error) {
	client := newApiClient()
	client.QueryParams["env_name"] = env
	client.QueryParams["force"] = fmt.Sprintf("%v", force)

	response := client.action(path.Join(serviceSetEntity, "undeploy", serviceSetName, "env", env)+"/", "DELETE", forceDeployServices)
	response.Process(true)

	var serviceResponse serviceset.ServiceSetServiceDeployResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *ServiceSet) UpdateServiceSet(serviceSetName string, serviceSet interface{}) {
	client := newApiClient()

	response := client.action(path.Join(serviceSetEntity, serviceSetName)+"/", "PUT", serviceSet)
	response.Process(true) // process response and exit if error
}
