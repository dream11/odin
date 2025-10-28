package auth

import (
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestGenerateState(t *testing.T) {
	t.Run("generates valid state with deterministic seed", func(t *testing.T) {
		seed := int64(12345)
		r := rand.New(rand.NewSource(seed))

		state, err := generateState(r)
		require.NoError(t, err)

		expectedState := "GulpVks0oz7NGvBf5pI9becYcJl9OO9gFVwyWVchTEI"
		assert.Equal(t, expectedState, state)

		assert.NotEmpty(t, state)
		for _, char := range state {
			assert.True(t, unicode.IsLetter(rune(char)) || unicode.IsDigit(rune(char)) ||
				char == '-' || char == '_',
				"generateState() contains invalid character: %c", char)

		}
		assert.Equal(t, 43, len(state), "generateState() returned state with unexpected length")
	})

	t.Run("generates unique states with different seeds", func(t *testing.T) {
		seed1 := int64(12345)
		seed2 := int64(67890)

		r1 := rand.New(rand.NewSource(seed1))
		r2 := rand.New(rand.NewSource(seed2))

		state1, err1 := generateState(r1)
		state2, err2 := generateState(r2)

		require.NoError(t, err1)
		require.NoError(t, err2)
		assert.NotEqual(t, state1, state2, "generateState() should generate unique states")
	})
}

func TestGetFreePort(t *testing.T) {
	port, err := getFreePort()
	require.NoError(t, err)
	assert.Greater(t, port, 0)
	assert.LessOrEqual(t, port, 65535)

	port2, err2 := getFreePort()
	require.NoError(t, err2)
	assert.Greater(t, port2, 0)
	assert.LessOrEqual(t, port2, 65535)
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
			name:        "nil input",
			input:       nil,
			wantErr:     true,
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
				assert.Equal(t, "test-client", config.ClientID)
				require.NotNil(t, config.AuthURL)
				assert.Equal(t, "https://example.com/auth", config.AuthURL.String())
				assert.Equal(t, "email", config.Scope)
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
				assert.Equal(t, "test-client-id", config.ClientID)
				assert.Equal(t, "openid profile email", config.Scope)
				assert.Equal(t, "test-provider", config.Name)
				require.NotNil(t, config.AuthURL)
				assert.Equal(t, "https://example.com/oauth/authorize", config.AuthURL.String())
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
				assert.Equal(t, "email", config.Scope)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := parseProviderData(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				return
			}
			require.NoError(t, err)
			require.NotNil(t, config)
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

	assert.Equal(t, "https://example.com/oauth/authorize", config.AuthURL.String(), "buildAuthURL() should not mutate original config.AuthURL")

	assert.Equal(t, "https", authURL.Scheme)
	assert.Equal(t, "example.com", authURL.Host)
	assert.Equal(t, "/oauth/authorize", authURL.Path)

	query := authURL.Query()
	assert.Equal(t, "test-client-id", query.Get("client_id"))
	assert.Equal(t, "code", query.Get("response_type"))
	assert.Equal(t, redirectURI, query.Get("redirect_uri"))
	assert.Equal(t, "openid profile email", query.Get("scope"))
	assert.Equal(t, state, query.Get("state"))
}

func TestSendSuccessPage(t *testing.T) {
	recorder := httptest.NewRecorder()
	sendSuccessPage(recorder)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "text/html; charset=utf-8", recorder.Header().Get("Content-Type"))

	body := recorder.Body.String()
	assert.Contains(t, body, "Authentication received")
	assert.Contains(t, body, "Completing sign-in in your terminal")
	assert.Contains(t, body, successPageTemplate)
}

func TestSendErrorPage(t *testing.T) {
	title := "Test Error"
	message := "This is a test error message"

	recorder := httptest.NewRecorder()
	sendErrorPage(recorder, title, message)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "text/html; charset=utf-8", recorder.Header().Get("Content-Type"))

	body := recorder.Body.String()
	assert.Contains(t, body, title)
	assert.Contains(t, body, message)
}

func mustStruct(m map[string]interface{}) *structpb.Struct {
	s, err := structpb.NewStruct(m)
	if err != nil {
		panic(err)
	}
	return s
}

func getFreePort() (int, error) {
	listener, err := net.Listen("tcp", net.JoinHostPort(localhost, "0"))
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}
