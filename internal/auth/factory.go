package auth

import (
	"fmt"
	"strings"
)

var providers = make(map[string]Provider)

func init() {
	providers["anonymous"] = &AnonymousProvider{}
	providers["oidc"] = &OIDCProvider{}
}

// GetProvider returns the requested authentication provider.
func GetProvider(providerType string) (Provider, error) {
	provider, ok := providers[strings.ToLower(providerType)]
	if !ok {
		return nil, fmt.Errorf("unknown authentication provider: %s", providerType)
	}
	return provider, nil
}
