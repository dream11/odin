package component

// Component interface
type Component struct {
	Name                      string      `yaml:"name,omitempty" json:"name,omitempty"`
	Type                      string      `yaml:"type,omitempty" json:"type,omitempty"`
	Version                   string      `yaml:"version,omitempty" json:"version,omitempty"`
	CreatedBy                 string      `yaml:"createdBy,omitempty" json:"createdBy,omitempty"`
	UpdatedBy                 string      `yaml:"updatedBy,omitempty" json:"updatedBy,omitempty"`
	CreatedAt                 string      `yaml:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt                 string      `yaml:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Active                    bool        `yaml:"isActive,omitempty" json:"isActive,omitempty"`
	Config                    interface{} `yaml:"config,omitempty" json:"config,omitempty"`
	Deployment                interface{} `yaml:"deployment_config,omitempty" json:"deployment_config,omitempty"`
	Scaling                   interface{} `yaml:"scaling_config,omitempty" json:"scaling_config,omitempty"`
	Discovery                 interface{} `yaml:"discovery_config,omitempty" json:"discovery_config,omitempty"`
	DeploymentPlatformMapping interface{} `yaml:"deployment_platform_mapping,omitempty" json:"deployment_platform_mapping,omitempty"`
}

// Type interface
type Type struct {
	Name                      string      `yaml:"name,omitempty" json:"name,omitempty"`
	Version                   string      `yaml:"version,omitempty" json:"version,omitempty"`
	CreatedBy                 string      `yaml:"createdBy,omitempty" json:"createdBy,omitempty"`
	UpdatedBy                 string      `yaml:"updatedBy,omitempty" json:"updatedBy,omitempty"`
	CreatedAt                 string      `yaml:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt                 string      `yaml:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Config                    interface{} `yaml:"config,omitempty" json:"config,omitempty"`
	Deployment                interface{} `yaml:"deployment_config,omitempty" json:"deployment_config,omitempty"`
	Scaling                   interface{} `yaml:"scaling_config,omitempty" json:"scaling_config,omitempty"`
	Discovery                 interface{} `yaml:"discovery_config,omitempty" json:"discovery_config,omitempty"`
	DeploymentPlatformMapping interface{} `yaml:"deployment_platform_mapping,omitempty" json:"deployment_platform_mapping,omitempty"`
}

// ListTypeResponse interface
type ListTypeResponse struct {
	Response []Type `yaml:"resp,omitempty" json:"resp,omitempty"`
}
