package deploy

import (
	v1 "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
)

func GetServiceDeployModel(response *v1.DeployServiceResponse) Model {
	serviceName := response.GetServiceResponse().Name

	serviceView := ServiceView{
		Name:           serviceName,
		Action:         response.GetServiceResponse().ServiceStatus.GetServiceAction(),
		Status:         response.GetServiceResponse().ServiceStatus.GetServiceStatus(),
		TraceId:        "random-trace-id",
		ComponentsView: make([]ComponentView, 0),
	}
	for _, component := range response.GetServiceResponse().GetComponentsStatus() {
		errorMessage := "Awaiting component logs"
		if component.GetError() != "" {
			errorMessage = component.GetError()
		}
		componentView := ComponentView{
			Name:    component.GetComponentName(),
			Action:  component.GetComponentAction(),
			Status:  component.GetComponentStatus(),
			Content: errorMessage,
		}
		serviceView.ComponentsView = append(serviceView.ComponentsView, componentView)
	}
	return Model{
		ServiceView: serviceView,
	}
}
