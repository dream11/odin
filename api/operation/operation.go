package operation

type Operation struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	InputSchema  interface{} `yaml:"inputSchema,omitempty" json:"inputSchema,omitempty"`
}

type ListOperation struct {
	Response []Operation `yaml:"resp,omitempty" json:"resp,omitempty"`
}

