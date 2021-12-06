package error

// Error interface
type Error struct {
	Error string `yaml:"err,omitempty" json:"err,omitempty"`
}
