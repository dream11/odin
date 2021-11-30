package component

// Artifact structure
type Artifact struct {
	URL     string `yaml:"url,omitempty" json:"url,omitempty"`
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
}

// Component structure
type Component struct {
	Name     string   `yaml:"name,omitempty" json:"name,omitempty"`
	Type     string   `yaml:"type,omitempty" json:"type,omitempty"`
	Artifact Artifact `yaml:"artifact,omitempty" json:"artifact,omitempty"`
}

// Type structure
type Type struct {
	Name        string      `yaml:"name,omitempty" json:"name,omitempty"`
	CreatedBy   string      `yaml:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedBy   string      `yaml:"updated_by,omitempty" json:"updated_by,omitempty"`
	CreatedAt   string      `yaml:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   string      `yaml:"updated_at,omitempty" json:"updated_at,omitempty"`
	Config      interface{} `yaml:"config,omitempty" json:"config,omitempty"`
	Deployment  interface{} `yaml:"deployment,omitempty" json:"deployment,omitempty"`
	Scaling     interface{} `yaml:"scaling,omitempty" json:"scaling,omitempty"`
	Discovery   interface{} `yaml:"discovery,omitempty" json:"discovery,omitempty"`
	MetaInfo    interface{} `yaml:"meta_info,omitempty" json:"meta_info,omitempty"`
	EnvBehavior interface{} `yaml:"env_specific_behavior,omitempty" json:"env_specific_behavior,omitempty"`
}

// ListTypeResponse structure
type ListTypeResponse struct {
	Response []Type `yaml:"resp,omitempty" json:"resp,omitempty"`
}
