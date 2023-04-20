package componenttype

// Component interface
type Component struct {
	Name                      string      `yaml:"name,omitempty" json:"name,omitempty"`
	Type                      string      `yaml:"type,omitempty" json:"type,omitempty"`
	Version                   string      `yaml:"version,omitempty" json:"version,omitempty"`
	Config                    interface{} `yaml:"config,omitempty" json:"config,omitempty"`
	Deployment                interface{} `yaml:"deployment_config,omitempty" json:"deployment_config,omitempty"`
	Scaling                   interface{} `yaml:"scaling_config,omitempty" json:"scaling_config,omitempty"`
	Discovery                 interface{} `yaml:"discovery_config,omitempty" json:"discovery_config,omitempty"`
	DeploymentPlatformMapping interface{} `yaml:"behaviour,omitempty" json:"behaviour,omitempty"`
}

// Type interface
type Type struct {
	Name                      string      `yaml:"name,omitempty" json:"name,omitempty"`
	Version                   string      `yaml:"version,omitempty" json:"version,omitempty"`
	TotalVersions             int         `yaml:"totalVersions,omitempty" json:"totalVersions,omitempty"`
	DeploymentTypes           []string    `yaml:"deployment_types,omitempty" json:"deployment_types,omitempty"`
	CreatedBy                 string      `yaml:"createdBy,omitempty" json:"createdBy,omitempty"`
	UpdatedBy                 string      `yaml:"updatedBy,omitempty" json:"updatedBy,omitempty"`
	CreatedAt                 string      `yaml:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt                 string      `yaml:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Config                    interface{} `yaml:"config,omitempty" json:"config,omitempty"`
	Deployment                interface{} `yaml:"deployment_config,omitempty" json:"deployment_config,omitempty"`
	Scaling                   interface{} `yaml:"scaling_config,omitempty" json:"scaling_config,omitempty"`
	Discovery                 interface{} `yaml:"discovery_config,omitempty" json:"discovery_config,omitempty"`
	DeploymentPlatformMapping interface{} `yaml:"behaviour,omitempty" json:"behaviour,omitempty"`
}

// Exposed Config interface
type ExposedConfig struct {
	Config      string `yaml:"config,omitempty" json:"config,omitempty"`
	Mandatory   bool   `yaml:"mandatory,omitempty" json:"mandatory,omitempty"`
	DataType    string `yaml:"data_type,omitempty" json:"data_type,omitempty"`
	Location    string `yaml:"location,omitempty" json:"location,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

// ListTypeResponse interface
type ListTypeResponse struct {
	Response []Type `yaml:"resp,omitempty" json:"resp,omitempty"`
}

// DetailComponentTypeResponse interface
type ComponentDetailsResponse struct {
	Response ComponentDetails `yaml:"resp,omitempty" json:"resp,omitempty"`
}

// ComponentDetails interface
type ComponentDetails struct {
	Details        Type            `yaml:"details,omitempty" json:"details,omitempty"`
	ExposedConfigs []ExposedConfig `yaml:"exposed_config,omitempty" json:"exposed_config,omitempty"`
}

// ListTypeResponse interface
type DetailComponentResponse struct {
	Response Component `yaml:"resp,omitempty" json:"resp,omitempty"`
}
