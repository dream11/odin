package environment

import (
	"github.com/dream11/odin/api/service"
)

// Env interface
type Env struct {
	Name         string            `yaml:"name,omitempty" json:"name,omitempty"`
	Team         string            `yaml:"team,omitempty" json:"team,omitempty"`
	Purpose      string            `yaml:"purpose,omitempty" json:"purpose,omitempty"`
	EnvType      string            `yaml:"env_type,omitempty" json:"env_type,omitempty"`
	State        string            `yaml:"state,omitempty" json:"state,omitempty"`
	DeletionTime string            `yaml:"deletion_time,omitempty" json:"deletion_time,omitempty"`
	Account      string            `yaml:"provider_account,omitempty" json:"provider_account,omitempty"`
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

type History struct {
	ID               int         `yaml:"id,omitempty" json:"id,omitempty"`
	CreatedBy        string      `yaml:"modifiedBy,omitempty" json:"createdBy,omitempty"`
	CreatedAt        string      `yaml:"lastModified,omitempty" json:"createdAt,omitempty"`
	EnvId            string      `yaml:"envName,omitempty" json:"envId,omitempty"`
	State            string      `yaml:"state,omitempty" json:"state,omitempty"`
	AutoDeletionTime string      `yaml:"autoDeletionTime,omitempty" json:"autoDeletionTime,omitempty"`
	EnvConfig        interface{} `yaml:"envConfig,omitempty" json:"envConfig,omitempty"`
}

// CreationResponse interface
type CreationResponse struct {
	Response Env `yaml:"resp,omitempty" json:"resp,omitempty"`
}

// ListResponse interface
type ListResponse struct {
	Response []Env `yaml:"resp,omitempty" json:"resp,omitempty"`
}

// HistoryListResponse interface
type HistoryListResponse struct {
	Response []History `yaml:"resp,omitempty" json:"resp,omitempty"`
}
