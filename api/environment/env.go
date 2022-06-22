package environment

import (
	"github.com/dream11/odin/api/service"
)

// Env interface
type Env struct {
	Name         string            `yaml:"name,omitempty" json:"name,omitempty"`
	Team         string            `yaml:"team,omitempty" json:"team,omitempty"`
	EnvType      string            `yaml:"envType,omitempty" json:"envType,omitempty"`
	State        string            `yaml:"state,omitempty" json:"state,omitempty"`
	DeletionTime string            `yaml:"autoDeletionTime,omitempty" json:"autoDeletionTime,omitempty"`
	Account      string            `yaml:"cloudProviderAccount,omitempty" json:"cloudProviderAccount,omitempty"`
	CreatedBy    string            `yaml:"createdBy,omitempty" json:"createdBy,omitempty"`
	UpdatedBy    string            `yaml:"updatedBy,omitempty" json:"updatedBy,omitempty"`
	CreatedAt    string            `yaml:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt    string            `yaml:"updatedAt,omitempty" json:"updatedAt,omitempty"`
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
	Action           string      `yaml:"action,omitempty" json:"action,omitempty"`
	ResourceDetails  string      `yaml:"resourceDetails,omitempty" json:"resourceDetails,omitempty"`
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

// EnvStatusResponse interface
type EnvStatusResponse struct {
	EnvResponse EnvStatus `yaml:"resp,omitempty" json:"resp,omitempty"`
}

type EnvDeleteResponse struct {
	EnvResponse EnvDelete `yaml:"resp,omitempty" json:"resp,omitempty"`
}

type EnvServiceStatusResponse struct {
	ServiceResponse EnvServiceStatus `yaml:"resp,omitempty" json:"resp,omitempty"`
}

type EnvTypesResponse struct {
	EnvTypes []string `yaml:"resp,omitempty" json:"resp,omitempty"`
}

type EnvDelete struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	ExecutorUrl string `yaml:"executorUrl,omitempty" json:"executorUrl,omitempty"`
}

type EnvServiceStatus struct {
	LastDeployedAt string   `yaml:"lastDeployedAt,omitempty" json:"lastDeployedAt,omitempty"`
	Version        string   `yaml:"version,omitempty" json:"version,omitempty"`
	Components     []Status `yaml:"components,omitempty" json:"components,omitempty"`
	Status         string   `yaml:"status,omitempty" json:"status,omitempty"`
}

type EnvStatus struct {
	ServiceStatus []EnvStatusPerService `yaml:"services,omitempty" json:"services,omitempty"`
	Status        string                `yaml:"status,omitempty" json:"status,omitempty"`
}

type EnvStatusPerService struct {
	Status         string `yaml:"status,omitempty" json:"status,omitempty"`
	Name           string `yaml:"name,omitempty" json:"name,omitempty"`
	Version        string `yaml:"version,omitempty" json:"version,omitempty"`
	LastDeployedAt string `yaml:"lastDeployedAt,omitempty" json:"lastDeployedAt,omitempty"`
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
