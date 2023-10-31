package configuration

// SecretKeys interface
type SecretKeys struct {
	AccessKey       string `toml:"access_key,omitempty" mapstructure:"access_key,omitempty"`
	SecretAccessKey string `toml:"secret_access_key,omitempty" mapstructure:"secret_access_key,omitempty"`
}

// Configuration interface
type Configuration struct {
	BackendAddr  string `toml:"backend_addr,omitempty" mapstructure:"backend_addr,omitempty"`
	Keys         SecretKeys
	AccessToken  string `toml:"access_token,omitempty" mapstructure:"access_token,omitempty"`
	EnvName      string `toml:"envName,omitempty" mapstructure:"envName,omitempty"`
	Insecure     bool   `toml:"insecure,omitempty" mapstructure:"insecure,omitempty"`
}
