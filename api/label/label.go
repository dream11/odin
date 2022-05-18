package label

type Label struct {
	Name        string                						  `yaml:"name,omitempty" json:"name,omitempty"`
	VersionCardinalityGreaterThanOne     bool                `yaml:"version_cardinality_greater_than_one,omitempty" json:"version_cardinality_greater_than_one,omitempty"`
}

type ListResponse struct {
	Response []Label `yaml:"resp,omitempty" json:"resp,omitempty"`
}
