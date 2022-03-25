package servicegroup

type Create struct {
	Message string `yaml:"message,omitempty" json:"message,omitempty"`
}

// CreateResponse struct
type CreateResponse struct {
	Response Create `yaml:"resp,omitempty" json:"resp,omitempty"`
}

type List struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
}

type ListResponse struct {
	Response []List `yaml:"resp,omitempty" json:"resp,omitempty"`
}
