package profile

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

type ProfileService struct {
	Name    string `yaml:"name,omitempty" json:"name,omitempty"`
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
}

type Describe struct {
	Name     string           `yaml:"name,omitempty" json:"name,omitempty"`
	Services []ProfileService `yaml:"services,omitempty" json:"services,omitempty"`
}

type DescribeResponse struct {
	Response Describe `yaml:"resp,omitempty" json:"resp,omitempty"`
}
