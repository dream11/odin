package infra

// Infra structure
type Infra struct {
	Name         string `yaml:"name,omitempty" json:"name,omitempty"`
	Team         string `yaml:"team,omitempty" json:"team,omitempty"`
	Reason       string `yaml:"reason,omitempty" json:"reason,omitempty"`
	Env          string `yaml:"env_type,omitempty" json:"env_type,omitempty"`
	State        string `yaml:"state,omitempty" json:"state,omitempty"`
	DeletionTime string `yaml:"deletion_time,omitempty" json:"deletion_time,omitempty"`
}

// CreationResponse structure
type CreationResponse struct {
	InfraName string `yaml:"infra_name,omitempty" json:"infra_name,omitempty"`
}
