package component

type artifact struct {
	URL     string `yaml:"url,omitempty" json:"url,omitempty"`
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
}

// Component interface
type Component struct {
	Name     string   `yaml:"name,omitempty" json:"name,omitempty"`
	Type     string   `yaml:"type,omitempty" json:"type,omitempty"`
	Artifact artifact `yaml:"artifact,omitempty" json:"artifact,omitempty"`
}

// Type interface
type Type struct {
	Name        string      `yaml:"name,omitempty" json:"name,omitempty"`
	CreatedBy   string      `yaml:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedBy   string      `yaml:"updated_by,omitempty" json:"updated_by,omitempty"`
	CreatedAt   string      `yaml:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   string      `yaml:"updated_at,omitempty" json:"updated_at,omitempty"`
	Config      interface{} `yaml:"config,omitempty" json:"config,omitempty"`
	Deployment  interface{} `yaml:"deployment_config,omitempty" json:"deployment_config,omitempty"`
	Scaling     interface{} `yaml:"scaling_config,omitempty" json:"scaling_config,omitempty"`
	Discovery   interface{} `yaml:"discovery_config,omitempty" json:"discovery_config,omitempty"`
	MetaInfo    interface{} `yaml:"meta_info,omitempty" json:"meta_info,omitempty"`
	EnvBehavior interface{} `yaml:"env_specific_behavior,omitempty" json:"env_specific_behavior,omitempty"`
}

// ListTypeResponse interface
type ListTypeResponse struct {
	Response []Type `yaml:"resp,omitempty" json:"resp,omitempty"`
}
