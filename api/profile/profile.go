package profile

type Service struct {
	Name       string    `yaml:"name" json:"name"`
	Version    string    `yaml:"version" json:"version"`
}

type Profile struct {
	Name        string       `yaml:"name" json:"name"`
	Version     string       `yaml:"version" json:"version"`
	Services    []Service    `yaml:"services" json:"services"`
}