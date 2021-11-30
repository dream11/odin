package service

// Service structure
type Service struct {
	Name        string   `yaml:"name,omitempty" json:"name,omitempty"`
	Version     string   `yaml:"version,omitempty" json:"version,omitempty"`
	Team        []string `yaml:"team,omitempty" json:"team,omitempty"`
	Description string   `yaml:"description,omitempty" json:"description,omitempty"`
	Mature      bool     `yaml:"isMature,omitempty" json:"isMature,omitempty"`
}

// ListResponse structure
type ListResponse struct {
	Response []Service `yaml:"resp,omitempty" json:"resp,omitempty"`
}
