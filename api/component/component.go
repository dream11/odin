package component

type Operation struct {
	Name   string      `yaml:"name,omitempty" json:"name,omitempty"`
	Values interface{} `yaml:"values,omitempty" json:"values,omitempty"`
}

type Data struct {
	EnvName     string      `yaml:"env_name,omitempty" json:"env_name,omitempty"`
	ServiceName string      `yaml:"service_name,omitempty" json:"service_name,omitempty"`
	Operations  []Operation `yaml:"operations,omitempty" json:"operations,omitempty"`
}

type OperateComponentRequest struct {
	Data Data `yaml:"data,omitempty" json:"data,omitempty"`
}

type CompareOperationChangesResponseBody struct {
	OldValues []map[string]interface{} `yaml:"old_values,omitempty" json:"old_values,omitempty"`
	NewValues []map[string]interface{} `yaml:"new_values,omitempty" json:"new_values,omitempty"`
}

type CompareOperationChangesResponse struct {
	Response CompareOperationChangesResponseBody `yaml:"resp,omitempty" json:"resp,omitempty"`
}
