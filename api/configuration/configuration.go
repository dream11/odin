package configuration

// SecretKeys interface
type SecretKeys struct {
	AccessKey       string `mapstructure:"access_key,omitempty"`
	SecretAccessKey string `mapstructure:"secret_access_key,omitempty"`
}

// Configuration interface
type Configuration struct {
	BackendAddr  string `mapstructure:"backend_addr,omitempty"`
	Keys         SecretKeys
	AccessToken  string `mapstructure:"access_token,omitempty"`
	RefreshToken string `mapstructure:"refresh_token,omitempty"`
	EnvName      string `mapstructure:"envName,omitempty"`
	Insecure     bool   `mapstructure:"insecure,omitempty"`
}
