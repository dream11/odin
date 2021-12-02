package auth

// Auth interface
type Auth struct {
	AccessToken  string `yaml:"access_token,omitempty" json:"access_token,omitempty"`
	RefreshToken string `yaml:"refresh_token,omitempty" json:"refresh_token,omitempty"`
}
