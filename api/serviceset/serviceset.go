package serviceset

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

type ServiceSetService struct {
	Name    string `yaml:"name,omitempty" json:"name,omitempty"`
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
}

type Describe struct {
	Name     string              `yaml:"name,omitempty" json:"name,omitempty"`
	Services []ServiceSetService `yaml:"services,omitempty" json:"services,omitempty"`
}

type DescribeResponse struct {
	Response Describe `yaml:"resp,omitempty" json:"resp,omitempty"`
}

type ListEnvService struct {
	Name       string `yaml:"name,omitempty" json:"name,omitempty"`
	Version    string `yaml:"version,omitempty" json:"version,omitempty"`
	EnvVersion string `yaml:"envVersion,omitempty" json:"envVersion,omitempty"`
}

type ListEnvServiceResponse struct {
	Response []ListEnvService `yaml:"resp,omitempty" json:"resp,omitempty"`
}

type ServiceSetServiceDeploy struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Version     string `yaml:"version,omitempty" json:"version,omitempty"`
	ExecutorUrl string `yaml:"envVersion,omitempty" json:"executorUrl,omitempty"`
	Error       string `yaml:"envVersion,omitempty" json:"error,omitempty"`
}

type ServiceSetServiceDeployResponse struct {
	Response []ServiceSetServiceDeploy `yaml:"resp,omitempty" json:"resp,omitempty"`
}
