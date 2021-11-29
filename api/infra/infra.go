package infra

// Infra structure
type Infra struct {
	Team         string `yaml:"team,omitempty" json:"team,omitempty"`
	Purpose      string `yaml:"purpose,omitempty" json:"purpose,omitempty"`
	Env          string `yaml:"env_type,omitempty" json:"env_type,omitempty"`
	State        string `yaml:"state,omitempty" json:"state,omitempty"`
	DeletionTime string `yaml:"deletion_time,omitempty" json:"deletion_time,omitempty"`
	Account      string `yaml:"provider_account,omitempty" json:"provider_account,omitempty"`
}

// CreationName structure
type CreationName struct {
	Name string `yaml:"infra_name,omitempty" json:"infra_name,omitempty"`
}

// CreationResponse structure
type CreationResponse struct {
	Response CreationName `yaml:"resp,omitempty" json:"resp,omitempty"`
}
