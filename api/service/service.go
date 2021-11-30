package service

import (
	"github.com/dream11/odin/api/component"
)

// Service structure
type Service struct {
	Name        string                `yaml:"name,omitempty" json:"name,omitempty"`
	Version     string                `yaml:"version,omitempty" json:"version,omitempty"`
	Team        []string              `yaml:"team,omitempty" json:"team,omitempty"`
	Description string                `yaml:"description,omitempty" json:"description,omitempty"`
	Mature      bool                  `yaml:"isMature,omitempty" json:"isMature,omitempty"`
	CreatedBy   string                `yaml:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedBy   string                `yaml:"updated_by,omitempty" json:"updated_by,omitempty"`
	CreatedAt   string                `yaml:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   string                `yaml:"updated_at,omitempty" json:"updated_at,omitempty"`
	Active      bool                  `yaml:"isActive,omitempty" json:"isActive,omitempty"`
	Tags        interface{}           `yaml:"tags,omitempty" json:"tags,omitempty"`
	MetaInfo    interface{}           `yaml:"meta_info,omitempty" json:"meta_info,omitempty"`
	Components  []component.Component `yaml:"components,omitempty" json:"components,omitempty"`
}

// ListResponse structure
type ListResponse struct {
	Response []Service `yaml:"resp,omitempty" json:"resp,omitempty"`
}
