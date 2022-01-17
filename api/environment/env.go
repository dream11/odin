package environment

import (
	"github.com/dream11/odin/api/service"
)

// Env interface
type Env struct {
	Name         string            `yaml:"name,omitempty" json:"name,omitempty"`
	Team         string            `yaml:"team,omitempty" json:"team,omitempty"`
	Purpose      string            `yaml:"purpose,omitempty" json:"purpose,omitempty"`
	EnvType      string            `yaml:"envType,omitempty" json:"envType,omitempty"`
	State        string            `yaml:"state,omitempty" json:"state,omitempty"`
	DeletionTime string            `yaml:"autoDeletionTime,omitempty" json:"autoDeletionTime,omitempty"`
	Account      string            `yaml:"cloudProviderAccount,omitempty" json:"cloudProviderAccount,omitempty"`
	CreatedBy    string            `yaml:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedBy    string            `yaml:"updated_by,omitempty" json:"updated_by,omitempty"`
	CreatedAt    string            `yaml:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt    string            `yaml:"updated_at,omitempty" json:"updated_at,omitempty"`
	Config       interface{}       `yaml:"config,omitempty" json:"config,omitempty"`
	MetaInfo     interface{}       `yaml:"meta_info,omitempty" json:"meta_info,omitempty"`
	Cost         string            `yaml:"cost,omitempty" json:"cost,omitempty"`
	Organization string            `yaml:"organization,omitempty" json:"organization,omitempty"`
	Services     []service.Service `yaml:"services,omitempty" json:"services,omitempty"`
}

// CreationResponse interface
type CreationResponse struct {
	Response Env `yaml:"resp,omitempty" json:"resp,omitempty"`
}

// ListResponse interface
type ListResponse struct {
	Response []Env `yaml:"resp,omitempty" json:"resp,omitempty"`
}

// EnvStatusResponse interface
type StatusResponse struct {
	Response EnvStatus `yaml:"resp,omitempty" json:"resp,omitempty"`
}

type EnvStatus struct {
	Status     string   `yaml:"status,omitempty" json:"status,omitempty"`
	Services   []Status `yaml:"services,omitempty" json:"services,omitempty"`
	Components []Status `yaml:"components,omitempty" json:"components,omitempty"`
}

type Status struct {
	Status  string `yaml:"status,omitempty" json:"status,omitempty"`
	Name    string `yaml:"name,omitempty" json:"name,omitempty"`
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
}

// DetailResponse interface
type DetailResponse struct {
	Response Env `yaml:"resp,omitempty" json:"resp,omitempty"`
}
