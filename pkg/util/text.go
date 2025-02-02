package util

import (
	"fmt"
	"github.com/dream11/odin/cmd/deploy"
	"strings"
)

func GetServiceHeader(s deploy.ServiceView) string {
	var header strings.Builder
	var action string
	switch s.Action {
	case "VALIDATE":
		action = "validating"
	case "DEPLOY":
		action = "deployment"
	case "UNDEPLOY":
		action = "un-deployment"
	}
	switch s.Status {
	case "IN_PROGRESS":
		header.WriteString(fmt.Sprintf("Service: %s  %s in progress", s.Name, action))
	case "SUCCESSFUL":
		header.WriteString(fmt.Sprintf("Service: %s  %s successful", s.Name, action))
	case "FAILED":
		header.WriteString(fmt.Sprintf("Service: %s  %s failed", s.Name, action))
	}
	return header.String()
}
