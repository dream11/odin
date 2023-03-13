package service

import (
	"github.com/dream11/odin/api/componenttype"
	"github.com/dream11/odin/api/label"
)

type Service struct {
	Name        string                    `yaml:"name,omitempty" json:"name,omitempty"`
	Version     string                    `yaml:"version,omitempty" json:"version,omitempty"`
	Team        string                    `yaml:"team,omitempty" json:"team,omitempty"`
	Description string                    `yaml:"description,omitempty" json:"description,omitempty"`
	CreatedBy   string                    `yaml:"createdBy,omitempty" json:"createdBy,omitempty"`
	UpdatedBy   string                    `yaml:"updatedBy,omitempty" json:"updatedBy,omitempty"`
	CreatedAt   string                    `yaml:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   string                    `yaml:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Active      *bool                     `yaml:"isActive,omitempty" json:"isActive,omitempty"`
	Tags        interface{}               `yaml:"tags,omitempty" json:"tags,omitempty"`
	Labels      []label.Label             `yaml:"labels,omitempty" json:"labels,omitempty"`
	Components  []componenttype.Component `yaml:"components,omitempty" json:"components,omitempty"`
}

// ListResponse interface
type ListResponse struct {
	Response []Service `yaml:"resp,omitempty" json:"resp,omitempty"`
}

// DetailResponse interface
type DetailResponse struct {
	Response Service `yaml:"resp,omitempty" json:"resp,omitempty"`
}

// Status interface
type Status struct {
	Name      string `yaml:"name,omitempty" json:"name,omitempty"`
	AWS_EC2   string `yaml:"aws_ec2,omitempty" json:"aws_ec2,omitempty"`
	Container string `yaml:"container,omitempty" json:"container,omitempty"`
}

// StatusResponse interface
type StatusResponse struct {
	Response []Status `yaml:"resp,omitempty" json:"resp,omitempty"`
}

// Service Release Body interface
type MergedService struct {
	Service            interface{}
	ProvisioningConfig map[string]interface{}
}

type Operation struct {
	Name   string      `yaml:"name,omitempty" json:"name,omitempty"`
	Data   interface{} `yaml:"data,omitempty" json:"data,omitempty"`
}

type OperationConsent struct {
	Name   string      `yaml:"name,omitempty" json:"name,omitempty"`
    IsFeedbackRequired bool `yaml:"is_feedback_required,omitempty" json:"is_feedback_required,omitempty"`
	Message string `yaml:"message,omitempty" json:"message,omitempty"`
}

type OperationValidationResponseBody struct {
	Operations []OperationConsent `yaml:"operations,omitempty" json:"operations,omitempty"`
}

// CompareResponse
type CompareResponse struct {
	Response interface{} `yaml:"resp,omitempty" json:"resp,omitempty"`
}

type OperationRequest struct {
	EnvName  string      `yaml:"env_name,omitempty" json:"env_name,omitempty"`
	Operations  []Operation `yaml:"operations,omitempty" json:"operations,omitempty"`
}

type OperationValidationResponse struct {
	Response OperationValidationResponseBody `yaml:"resp,omitempty" json:"resp,omitempty"`
}
