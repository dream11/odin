package auth

import "google.golang.org/protobuf/types/known/structpb"

type AnonymousProvider struct{}

func (p *AnonymousProvider) Authenticate(providerData *structpb.Struct) (*structpb.Struct, error) {
	return &structpb.Struct{}, nil
}
