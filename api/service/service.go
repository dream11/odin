package service

import (
	"github.com/dream11/odin/api/component"
	"github.com/dream11/odin/api/label"
)

type Service struct {
	Name        string                `yaml:"name,omitempty" json:"name,omitempty"`
	Version     string                `yaml:"version,omitempty" json:"version,omitempty"`
	Team        []string              `yaml:"team,omitempty" json:"team,omitempty"`
	Description string                `yaml:"description,omitempty" json:"description,omitempty"`
	Mature      *bool                 `yaml:"isMature,omitempty" json:"isMature,omitempty"`
	CreatedBy   string                `yaml:"createdBy,omitempty" json:"createdBy,omitempty"`
	UpdatedBy   string                `yaml:"updatedBy,omitempty" json:"updatedBy,omitempty"`
	CreatedAt   string                `yaml:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   string                `yaml:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Active      *bool                 `yaml:"isActive,omitempty" json:"isActive,omitempty"`
	Tags        interface{}           `yaml:"tags,omitempty" json:"tags,omitempty"`
	Labels      []label.Label         `yaml:"labels,omitempty" json:"labels,omitempty"`
	Components  []component.Component `yaml:"components,omitempty" json:"components,omitempty"`
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
	VM        string `yaml:"vm,omitempty" json:"vm,omitempty"`
	Container string `yaml:"container,omitempty" json:"container,omitempty"`
}

// StatusResponse interface
type StatusResponse struct {
	Response []Status `yaml:"resp,omitempty" json:"resp,omitempty"`
}
