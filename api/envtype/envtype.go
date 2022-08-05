package envtype

// ListResponse interface
type ListTypeResponse struct {
	Response []Type
}

type Type struct {
	Things []string
}
