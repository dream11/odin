package infra

// Infra structure
type Infra struct {
	Name         string `yaml:"name,omitempty" json:"name,omitempty"`
	Team         string `yaml:"team,omitempty" json:"team,omitempty"`
	Purpose      string `yaml:"purpose,omitempty" json:"purpose,omitempty"`
	Env          string `yaml:"env_type,omitempty" json:"env_type,omitempty"`
	State        string `yaml:"state,omitempty" json:"state,omitempty"`
	DeletionTime string `yaml:"deletion_time,omitempty" json:"deletion_time,omitempty"`
	Account      string `yaml:"provider_account,omitempty" json:"provider_account,omitempty"`
	CreatedBy    string `yaml:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedBy    string `yaml:"updated_by,omitempty" json:"updated_by,omitempty"`
	CreatedAt    string `yaml:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt    string `yaml:"updated_at,omitempty" json:"updated_at,omitempty"`
	Config       interface{} `yaml:"config,omitempty" json:"config,omitempty"`
	MetaInfo     interface{} `yaml:"meta_info,omitempty" json:"meta_info,omitempty"`
	Cost         string `yaml:"cost,omitempty" json:"cost,omitempty"`
	Organization string `yaml:"organization,omitempty" json:"organization,omitempty"`
	Services     []string `yaml:"services,omitempty" json:"services,omitempty"`
}

// CreationName structure
// TODO: to be deprecated by Infra struct
type CreationName struct {
	Name string `yaml:"infra_name,omitempty" json:"infra_name,omitempty"`
}

// CreationResponse structure
type CreationResponse struct {
	Response CreationName `yaml:"resp,omitempty" json:"resp,omitempty"`
}

// ListResponse structure
type ListResponse struct {
	Response []Infra `yaml:"resp,omitempty" json:"resp,omitempty"`
}
