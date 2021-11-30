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
