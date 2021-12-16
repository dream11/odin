package component

// Component interface
type Component struct {
	Name        string      `yaml:"name,omitempty" json:"name,omitempty"`
	Type        string      `yaml:"type,omitempty" json:"type,omitempty"`
	Version     string      `yaml:"version,omitempty" json:"version,omitempty"`
	CreatedBy   string      `yaml:"createdBy,omitempty" json:"createdBy,omitempty"`
	UpdatedBy   string      `yaml:"updatedBy,omitempty" json:"updatedBy,omitempty"`
	CreatedAt   string      `yaml:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   string      `yaml:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Active      bool        `yaml:"isActive,omitempty" json:"isActive,omitempty"`
	Config      interface{} `yaml:"config,omitempty" json:"config,omitempty"`
	Deployment  interface{} `yaml:"deploymentConfig,omitempty" json:"deploymentConfig,omitempty"`
	Scaling     interface{} `yaml:"scalingConfig,omitempty" json:"scalingConfig,omitempty"`
	Discovery   interface{} `yaml:"discoveryConfig,omitempty" json:"discoveryConfig,omitempty"`
	EnvBehavior interface{} `yaml:"deploymentPlatformMapping,omitempty" json:"deploymentPlatformMapping,omitempty"`
}

// Type interface
type Type struct {
	Name        string      `yaml:"name,omitempty" json:"name,omitempty"`
	CreatedBy   string      `yaml:"createdBy,omitempty" json:"createdBy,omitempty"`
	UpdatedBy   string      `yaml:"updatedBy,omitempty" json:"updatedBy,omitempty"`
	CreatedAt   string      `yaml:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   string      `yaml:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Config      interface{} `yaml:"config,omitempty" json:"config,omitempty"`
	Deployment  interface{} `yaml:"deploymentConfig,omitempty" json:"deploymentConfig,omitempty"`
	Scaling     interface{} `yaml:"scalingConfig,omitempty" json:"scalingConfig,omitempty"`
	Discovery   interface{} `yaml:"discoveryConfig,omitempty" json:"discovery_discoveryConfigconfig,omitempty"`
	EnvBehavior interface{} `yaml:"deploymentPlatformMapping,omitempty" json:"deploymentPlatformMapping,omitempty"`
}

// ListTypeResponse interface
type ListTypeResponse struct {
	Response []Type `yaml:"resp,omitempty" json:"resp,omitempty"`
}
