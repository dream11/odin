package configuration

// Configuration interface
type Configuration struct {
	BackendAddress string `toml:"backend_address,omitempty" mapstructure:"backend_address,omitempty"`
	AccessToken    string `toml:"access_token,omitempty" mapstructure:"access_token,omitempty"`
	OrgId          int64  `toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	EnvName        string `toml:"envName,omitempty" mapstructure:"envName,omitempty"`
	Insecure       bool   `toml:"insecure,omitempty" mapstructure:"insecure,omitempty"`
	Plaintext      bool   `toml:"plaintext,omitempty" mapstructure:"plaintext,omitempty"`
}
