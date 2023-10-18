package configuration

// SecretKeys interface
type SecretKeys struct {
	AccessKey       string `yaml:"access_key,omitempty" json:"access_key,omitempty"`
	SecretAccessKey string `yaml:"secret_access_key,omitempty" json:"secret_access_key,omitempty"`
}

// Configuration interface
type Configuration struct {
	BackendAddr  string `yaml:"backend_addr,omitempty" json:"backend_addr,omitempty"`
	Keys         SecretKeys
	EnvName      string `yaml:"envName,omitempty" json:"envName:,omitempty"`
}
