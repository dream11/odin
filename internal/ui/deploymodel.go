package ui

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	v1 "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
)

type ServiceDeployModel struct {
	ServiceDisplayMeta ServiceDisplayMeta
	ServiceView        ServiceView
}

type ServiceDisplayMeta struct {
	Height               int
	Width                int
	Ready                bool
	Cursor               int
	ElapsedTime          int
	TotalCompletionTime  int
	Progress             progress.Model
	ComponentDisplayMeta []ComponentDisplayMeta
}

type ComponentDisplayMeta struct {
	Height      int
	Width       int
	Toggle      bool
	Spinner     spinner.Model
	LogViewPort viewport.Model
}

type ServiceView struct {
	Name           string
	Status         string
	Action         string
	TraceID        string
	ComponentsView []ComponentView
}

type ComponentView struct {
	Name    string
	Status  string
	Action  string
	Content string
}

func GetServiceDeployModel(response *v1.DeployServiceResponse, traceID string) ServiceDeployModel {
	serviceName := response.GetServiceResponse().Name

	serviceView := ServiceView{
		Name:           serviceName,
		Action:         response.GetServiceResponse().ServiceStatus.GetServiceAction(),
		Status:         response.GetServiceResponse().ServiceStatus.GetServiceStatus(),
		TraceID:        traceID,
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
	return ServiceDeployModel{
		ServiceView: serviceView,
	}
}
