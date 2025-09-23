package auth

import "google.golang.org/protobuf/types/known/structpb"

// Provider Interface that all authentication methods must implement.
type Provider interface {
	Authenticate(providerData *structpb.Struct) (*structpb.Struct, error)
}
