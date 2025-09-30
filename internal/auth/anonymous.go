package auth

import "google.golang.org/protobuf/types/known/structpb"

// AnonymousProvider implements a no-op authentication provider.
type AnonymousProvider struct{}

// Authenticate returns an empty struct without performing authentication.
func (p *AnonymousProvider) Authenticate(providerData *structpb.Struct) (*structpb.Struct, error) {
	return &structpb.Struct{}, nil
}
