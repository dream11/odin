package deploy

import (
	"fmt"
	v1 "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
)

func GetServiceView(response *v1.DeployServiceResponse) ServiceView {
	serviceName := response.GetServiceResponse().Name

	serviceHeaderText := fmt.Sprintf("Deploying service %s ", serviceName)
	if response.GetServiceResponse().ServiceStatus.GetServiceStatus() == "SUCCESSFUL" {
		serviceHeaderText = fmt.Sprintf("Service %s deployed successfully", serviceName)
	} else if response.GetServiceResponse().ServiceStatus.GetServiceStatus() == "FAILED" {
		serviceHeaderText = fmt.Sprintf("Failed to deploy service %s", serviceName)
	}

	serviceView := ServiceView{
		Header: Header{
			Text: serviceHeaderText,
		},
		Status:         response.GetServiceResponse().ServiceStatus.GetServiceStatus(),
		ComponentsView: make([]ComponentView, 0),
	}
	for _, component := range response.GetServiceResponse().GetComponentsStatus() {
		componentHeaderText := fmt.Sprintf("Deploying component %s", component.GetComponentName())
		if component.GetComponentStatus() == "SUCCESSFUL" {
			componentHeaderText = fmt.Sprintf("Component %s deployed successfully", component.GetComponentName())
		} else if component.GetComponentStatus() == "FAILED" {
			componentHeaderText = fmt.Sprintf("Failed to deploy component %s", component.GetComponentName())
		}

		errorMessage := "Awaiting component logs"
		if component.GetError() != "" {
			errorMessage = component.GetError()
		}
		componentView := ComponentView{
			Header: Header{Text: componentHeaderText},
			Status: component.GetComponentStatus(),
			LogView: LogView{
				Content: errorMessage,
			},
		}
		serviceView.ComponentsView = append(serviceView.ComponentsView, componentView)
	}
	return serviceView
}
