package backend

import (
	"encoding/json"
	"fmt"
	"path"
	"strconv"

	"github.com/dream11/odin/api/serviceset"
	"github.com/dream11/odin/pkg/request"
)

// ServiceSet entity
type ServiceSet struct{}

// root entity
var serviceSetEntity = "serviceset"
var serviceSetEntityFile = "servicesetfile"

func (s *ServiceSet) CreateServiceSet(serviceSetDefinition interface{}) {
	client := newApiClient()

	response := client.actionWithRetry(serviceSetEntity+"/", "POST", serviceSetDefinition)
	response.ProcessHandleError(true)
}

func (s *ServiceSet) CreateTempServiceSet(serviceSetDefinition interface{}) {
	client := newApiClient()

	response := client.actionWithRetry(serviceSetEntityFile+"/", "POST", serviceSetDefinition)
	response.ProcessHandleError(true)
}

func (s *ServiceSet) ListServiceSet(serviceSetName, serviceName string) ([]serviceset.List, error) {
	client := newApiClient()
	client.QueryParams["name"] = serviceSetName
	client.QueryParams["service"] = serviceName

	response := client.actionWithRetry(serviceSetEntity, "GET", nil)
	response.Process(true)

	var serviceResponse serviceset.ListResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *ServiceSet) DescribeServiceSet(serviceSetName string, isFile bool) (serviceset.Describe, error) {
	client := newApiClient()
	var response request.Response
	if isFile {
		response = client.actionWithRetry(path.Join(serviceSetEntityFile, serviceSetName), "GET", nil)
	} else {
		response = client.actionWithRetry(path.Join(serviceSetEntity, serviceSetName), "GET", nil)
	}
	response.Process(true)
	var serviceResponse serviceset.DescribeResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *ServiceSet) IdentifyServiceSetType(serviceSetName string) bool {
	isFile := false
	client := newApiClient()
	response := client.actionWithRetry(path.Join(serviceSetEntityFile, serviceSetName), "GET", nil)

	if response.GetStatusCode(true) == 200 {
		isFile = true
	} else if response.GetStatusCode(true) == 500 {
		response.Process(true)
	}
	return isFile
}

func (s *ServiceSet) DeleteServiceSet(serviceSetName string, isFile bool) {

	client := newApiClient()
	var response request.Response
	if isFile {
		response = client.actionWithRetry(path.Join(serviceSetEntityFile, serviceSetName)+"/", "DELETE", nil)

	} else {
		response = client.actionWithRetry(path.Join(serviceSetEntity, serviceSetName)+"/", "DELETE", nil)
	}

	response.Process(true)
}

func (s *ServiceSet) DeployServiceSet(serviceSetName, env, configStoreNamespace string, forceDeployServices []serviceset.ListEnvService, force bool, isFile bool) {
	client := newStreamingApiClient()
	client.QueryParams["env_name"] = env
	client.QueryParams["force"] = fmt.Sprintf("%v", force)
	client.QueryParams["config_store_namespace"] = configStoreNamespace
	client.QueryParams["isFile"] = fmt.Sprintf("%v", isFile)

	data := map[string]interface{}{
		"forceDeployServices": forceDeployServices,
	}

	response := client.streamWithRetry(path.Join(serviceSetEntity, "deploy", serviceSetName, "env", env)+"/", "POST", data)
	response.Process(true)

}

func (s ServiceSet) ListEnvServices(serviceSetName, env, filterBy string, isFile bool) ([]serviceset.ListEnvService, error) {
	client := newApiClient()
	client.QueryParams["filter_by"] = filterBy
	client.QueryParams["isFile"] = strconv.FormatBool(isFile)

	response := client.actionWithRetry(path.Join(serviceSetEntity, serviceSetName, "env", env, "service")+"/", "GET", nil)
	response.Process(true)

	var serviceResponse serviceset.ListEnvServiceResponse
	err := json.Unmarshal(response.Body, &serviceResponse)

	return serviceResponse.Response, err
}

func (s *ServiceSet) UndeployServiceSet(serviceSetName, env string, forceUndeployServices []serviceset.ListEnvService, force bool, isFile bool) {
	client := newStreamingApiClient()
	client.QueryParams["env_name"] = env
	client.QueryParams["force"] = fmt.Sprintf("%v", force)
	client.QueryParams["isFile"] = fmt.Sprintf("%v", isFile)

	data := map[string]interface{}{
		"forceUndeployServices": forceUndeployServices,
	}

	response := client.streamWithRetry(path.Join(serviceSetEntity, "undeploy", serviceSetName, "env", env)+"/", "DELETE", data)
	response.Process(true)
}

func (s *ServiceSet) UpdateServiceSet(serviceSetName string, serviceSet interface{}) {
	client := newApiClient()

	response := client.actionWithRetry(path.Join(serviceSetEntity, serviceSetName)+"/", "PUT", serviceSet)
	response.Process(true) // process response and exit if error
}
