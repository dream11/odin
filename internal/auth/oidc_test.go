package auth

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"google.golang.org/protobuf/types/known/structpb"
)

func TestGenerateState(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "generates valid state",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state, err := generateState()
			if (err != nil) != tt.wantErr {
				t.Errorf("generateState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if state == "" {
					t.Error("generateState() returned empty state")
				}
				for _, char := range state {
					if !((char >= 'A' && char <= 'Z') ||
						(char >= 'a' && char <= 'z') ||
						(char >= '0' && char <= '9') ||
						char == '-' || char == '_') {
						t.Errorf("generateState() contains invalid character: %c", char)
					}
				}
				if len(state) < 40 || len(state) > 44 {
					t.Errorf("generateState() returned state with unexpected length: %d (expected ~43)", len(state))
				}
			}
		})
	}

	state1, err1 := generateState()
	state2, err2 := generateState()
	if err1 != nil || err2 != nil {
		t.Fatal("generateState() should not return errors")
	}
	if state1 == state2 {
		t.Error("generateState() should generate unique states, but got identical values")
	}
}

func TestGetFreePort(t *testing.T) {
	port, err := getFreePort()
	if err != nil {
		t.Fatalf("getFreePort() error = %v", err)
	}
	if port <= 0 || port > 65535 {
		t.Errorf("getFreePort() returned invalid port: %d (expected 1-65535)", port)
	}

	port2, err2 := getFreePort()
	if err2 != nil {
		t.Fatalf("getFreePort() second call error = %v", err2)
	}
	if port2 <= 0 || port2 > 65535 {
		t.Errorf("getFreePort() second call returned invalid port: %d", port2)
	}
}

func TestParseProviderData(t *testing.T) {
	tests := []struct {
		name        string
		input       *structpb.Struct
		wantErr     bool
		errContains string
		validate    func(*testing.T, *OIDCProviderConfig)
	}{
		{
			name:    "nil input",
			input:   nil,
			wantErr: true,
			errContains: "provider data is required",
		},
		{
			name: "missing authorization_url",
			input: mustStruct(map[string]interface{}{
				"client_id": "test-client",
			}),
			wantErr:     true,
			errContains: "authorization_url is required",
		},
		{
			name: "missing client_id",
			input: mustStruct(map[string]interface{}{
				"authorization_url": "https://example.com/auth",
			}),
			wantErr:     true,
			errContains: "client_id is required",
		},
		{
			name: "invalid authorization_url",
			input: mustStruct(map[string]interface{}{
				"authorization_url": "://invalid-url",
				"client_id":         "test-client",
			}),
			wantErr:     true,
			errContains: "invalid authorization_url",
		},
		{
			name: "valid minimal config",
			input: mustStruct(map[string]interface{}{
				"authorization_url": "https://example.com/auth",
				"client_id":         "test-client",
			}),
			wantErr: false,
			validate: func(t *testing.T, config *OIDCProviderConfig) {
				if config.ClientID != "test-client" {
					t.Errorf("expected client_id 'test-client', got %q", config.ClientID)
				}
				if config.AuthURL == nil || config.AuthURL.String() != "https://example.com/auth" {
					t.Errorf("expected auth URL 'https://example.com/auth', got %v", config.AuthURL)
				}
				if config.Scope != "email" {
					t.Errorf("expected default scope 'email', got %q", config.Scope)
				}
			},
		},
		{
			name: "valid config with all fields",
			input: mustStruct(map[string]interface{}{
				"authorization_url": "https://example.com/oauth/authorize",
				"client_id":         "test-client-id",
				"scope":             "openid profile email",
				"name":              "test-provider",
			}),
			wantErr: false,
			validate: func(t *testing.T, config *OIDCProviderConfig) {
				if config.ClientID != "test-client-id" {
					t.Errorf("expected client_id 'test-client-id', got %q", config.ClientID)
				}
				if config.Scope != "openid profile email" {
					t.Errorf("expected scope 'openid profile email', got %q", config.Scope)
				}
				if config.Name != "test-provider" {
					t.Errorf("expected name 'test-provider', got %q", config.Name)
				}
				if config.AuthURL == nil || config.AuthURL.String() != "https://example.com/oauth/authorize" {
					t.Errorf("expected auth URL 'https://example.com/oauth/authorize', got %v", config.AuthURL)
				}
			},
		},
		{
			name: "empty scope defaults to email",
			input: mustStruct(map[string]interface{}{
				"authorization_url": "https://example.com/auth",
				"client_id":         "test-client",
				"scope":             "",
			}),
			wantErr: false,
			validate: func(t *testing.T, config *OIDCProviderConfig) {
				if config.Scope != "email" {
					t.Errorf("expected default scope 'email', got %q", config.Scope)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := parseProviderData(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseProviderData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("parseProviderData() error = %v, should contain %q", err, tt.errContains)
				}
				return
			}
			if config == nil {
				t.Fatal("parseProviderData() returned nil config without error")
			}
			if tt.validate != nil {
				tt.validate(t, config)
			}
		})
	}
}

func TestBuildAuthURL(t *testing.T) {
	baseURL, _ := url.Parse("https://example.com/oauth/authorize")
	config := &OIDCProviderConfig{
		AuthURL:  baseURL,
		ClientID: "test-client-id",
		Scope:    "openid profile email",
	}

	redirectURI := "http://localhost:8080/callback"
	state := "test-state-123"

	authURL := buildAuthURL(config, redirectURI, state)

	if config.AuthURL.String() != "https://example.com/oauth/authorize" {
		t.Errorf("buildAuthURL() mutated original config.AuthURL: %v", config.AuthURL)
	}

	if authURL.Scheme != "https" || authURL.Host != "example.com" || authURL.Path != "/oauth/authorize" {
		t.Errorf("buildAuthURL() incorrect base URL: %v", authURL)
	}

	query := authURL.Query()
	if query.Get("client_id") != "test-client-id" {
		t.Errorf("buildAuthURL() missing or incorrect client_id: %q", query.Get("client_id"))
	}
	if query.Get("response_type") != "code" {
		t.Errorf("buildAuthURL() missing or incorrect response_type: %q", query.Get("response_type"))
	}
	if query.Get("redirect_uri") != redirectURI {
		t.Errorf("buildAuthURL() missing or incorrect redirect_uri: %q", query.Get("redirect_uri"))
	}
	if query.Get("scope") != "openid profile email" {
		t.Errorf("buildAuthURL() missing or incorrect scope: %q", query.Get("scope"))
	}
	if query.Get("state") != state {
		t.Errorf("buildAuthURL() missing or incorrect state: %q", query.Get("state"))
	}
}

func TestSendSuccessPage(t *testing.T) {
	recorder := httptest.NewRecorder()
	sendSuccessPage(recorder)

	if recorder.Code != http.StatusOK {
		t.Errorf("sendSuccessPage() status code = %d, want %d", recorder.Code, http.StatusOK)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "text/html; charset=utf-8" {
		t.Errorf("sendSuccessPage() Content-Type = %q, want %q", contentType, "text/html; charset=utf-8")
	}

	body := recorder.Body.String()
	if !strings.Contains(body, "Authentication received") {
		t.Error("sendSuccessPage() body should contain 'Authentication received'")
	}
	if !strings.Contains(body, "Completing sign-in in your terminal") {
		t.Error("sendSuccessPage() body should contain 'Completing sign-in in your terminal'")
	}
	if !strings.Contains(body, successPageTemplate) {
		t.Error("sendSuccessPage() body should match successPageTemplate")
	}
}

func TestSendErrorPage(t *testing.T) {
	title := "Test Error"
	message := "This is a test error message"

	recorder := httptest.NewRecorder()
	sendErrorPage(recorder, title, message)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("sendErrorPage() status code = %d, want %d", recorder.Code, http.StatusBadRequest)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "text/html; charset=utf-8" {
		t.Errorf("sendErrorPage() Content-Type = %q, want %q", contentType, "text/html; charset=utf-8")
	}

	body := recorder.Body.String()
	if !strings.Contains(body, title) {
		t.Errorf("sendErrorPage() body should contain title %q", title)
	}
	if !strings.Contains(body, message) {
		t.Errorf("sendErrorPage() body should contain message %q", message)
	}
}

func mustStruct(m map[string]interface{}) *structpb.Struct {
	s, err := structpb.NewStruct(m)
	if err != nil {
		panic(err)
	}
	return s
}

