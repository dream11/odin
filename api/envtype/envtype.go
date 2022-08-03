package envtype

// Env Type interface
type EnvType struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
}

// ListResponse interface
type ListResponse struct {
	Response []EnvType `yaml:"resp,omitempty" json:"resp,omitempty"`
}
