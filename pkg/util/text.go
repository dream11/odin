package util

import (
	"fmt"
	"strings"
)

func GetHeaderText(name string, action string, status string, element string) string {
	var header strings.Builder
	var actionText string
	switch action {
	case "VALIDATE":
		actionText = "validating"
	case "DEPLOY":
		actionText = "deployment"
	case "UNDEPLOY":
		actionText = "un-deployment"
	}
	switch status {
	case "IN_PROGRESS":
		header.WriteString(fmt.Sprintf("%s: %s  %s in progress", element, name, actionText))
	case "SUCCESSFUL":
		header.WriteString(fmt.Sprintf("%s: %s  %s successful", element, name, actionText))
	case "FAILED":
		header.WriteString(fmt.Sprintf("%s: %s  %s failed", element, name, actionText))
	}
	return header.String()
}
