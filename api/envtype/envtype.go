package envtype

type EnvType struct {
	Id         	 int64             `yaml:"id,omitempty" json:"id,omitempty"`
	CreatedBy    string            `yaml:"createdBy,omitempty" json:"createdBy,omitempty"`
	UpdatedBy    string            `yaml:"updatedBy,omitempty" json:"updatedBy,omitempty"`
	CreatedAt    string            `yaml:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt    string            `yaml:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Name		 string			   `yaml:"name,omitempty" json:"name,omitempty"`
	Strict       bool			   `yaml:"strict,omitempty" json:"strict,omitempty"`
}

// ListResponse interface
type ListTypeResponse struct {
	Response []string `yaml:"resp,omitempty" json:"resp,omitempty"`
}

type GetEnvTypeResponse struct {
	Response EnvType `yaml:"resp,omitempty" json:"resp,omitempty"`
}
