package auth

import (
	"context"
	"crypto/rand"
	_ "embed"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
)

//go:embed templates/success.html
var successPageTemplate string

//go:embed templates/error.html
var errorPageTemplate string

const (
	callbackPath               = "/callback"
	redirectScheme             = "http"
	localhost                  = "localhost"
	readHeaderTimeout          = 5 * time.Second
	closeAfterWriteDelay       = 200 * time.Millisecond
	defaultCallbackWaitTimeout = 5 * time.Minute
)

type OIDCProviderConfig struct {
	Name     string
	AuthURL  *url.URL
	ClientID string
	Scope    string
}

type OIDCProvider struct{}

func (p *OIDCProvider) Authenticate(providerData *structpb.Struct) (*structpb.Struct, error) {
	config, err := parseProviderData(providerData)
	if err != nil {
		return nil, fmt.Errorf("parse provider data: %w", err)
	}

	ln, err := net.Listen("tcp", net.JoinHostPort(localhost, "0"))
	if err != nil {
		return nil, fmt.Errorf("listen on callback port: %w", err)
	}
	port := ln.Addr().(*net.TCPAddr).Port

	redirectURL := &url.URL{
		Scheme: redirectScheme,
		Host:   net.JoinHostPort(localhost, fmt.Sprintf("%d", port)),
		Path:   callbackPath,
	}

	state, err := generateState(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("generate state: %w", err)
	}

	authURL := buildAuthURL(config, redirectURL.String(), state)

	if err := openBrowser(authURL); err != nil {
		log.Warnf("Failed to open browser automatically: %v", err)
	}

	log.Info("\nPlease visit the following URL to authenticate:")
	log.Info(authURL.String())

	authCode, err := waitForCallback(ln, state)
	if err != nil {
		return nil, err
	}

	return structpb.NewStruct(map[string]interface{}{
		"authorization_code": authCode,
		"redirect_uri":       redirectURL.String(),
	})
}

func waitForCallback(ln net.Listener, expectedState string) (string, error) {
	defer ln.Close()

	var code string

	srv := &http.Server{
		ReadHeaderTimeout: readHeaderTimeout,
	}

	sendErrorAndShutdown := func(w http.ResponseWriter, srv *http.Server, title, message string) {
		sendErrorPage(w, title, message)
		ctx, cancel := context.WithTimeout(context.Background(), closeAfterWriteDelay)
		time.AfterFunc(closeAfterWriteDelay, func() {
			_ = srv.Shutdown(ctx)
			cancel()
		})
	}

	sendSuccessAndShutdown := func(w http.ResponseWriter, srv *http.Server) {
		sendSuccessPage(w)
		ctx, cancel := context.WithTimeout(context.Background(), closeAfterWriteDelay)
		time.AfterFunc(closeAfterWriteDelay, func() {
			_ = srv.Shutdown(ctx)
			cancel()
		})
	}

	srv.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != callbackPath {
			http.NotFound(w, r)
			return
		}
		q := r.URL.Query()

		if e := q.Get("error"); e != "" {
			sendErrorAndShutdown(w, srv, "Authentication Failed", q.Get("error_description"))
			return
		}

		if q.Get("state") != expectedState {
			sendErrorAndShutdown(w, srv, "Security Error", "Invalid state parameter")
			return
		}

		c := q.Get("code")
		if c == "" {
			sendErrorAndShutdown(w, srv, "Authentication Failed", "No authorization code received")
			return
		}

		code = c
		sendSuccessAndShutdown(w, srv)
	})

	ctx, cancel := context.WithTimeout(context.Background(), defaultCallbackWaitTimeout)
	t := time.AfterFunc(defaultCallbackWaitTimeout, func() {
		_ = srv.Shutdown(ctx)
		cancel()
	})

	defer cancel()
	defer t.Stop()

	if err := srv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return "", fmt.Errorf("serve error: %w", err)
	}

	if code == "" {
		return "", fmt.Errorf("authentication failed or timed out after %s", defaultCallbackWaitTimeout)
	}
	return code, nil
}

func sendSuccessPage(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := fmt.Fprint(w, successPageTemplate); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
	}
}

func sendErrorPage(w http.ResponseWriter, title, message string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, errorPageTemplate, title, title, message)
}

func parseProviderData(data *structpb.Struct) (*OIDCProviderConfig, error) {
	if data == nil {
		return nil, fmt.Errorf("provider data is required")
	}

	fields := data.GetFields()

	getString := func(m map[string]*structpb.Value, key string, defaultValue string) string {
		if v, ok := m[key]; ok && v != nil {
			val := v.GetStringValue()
			if val != "" {
				return val
			}
		}
		return defaultValue
	}

	config := &OIDCProviderConfig{
		Name:     getString(fields, "name", ""),
		ClientID: getString(fields, "client_id", ""),
		Scope:    getString(fields, "scope", "email"),
	}

	authURLStr := getString(fields, "authorization_url", "")
	if authURLStr == "" {
		return nil, fmt.Errorf("authorization_url is required")
	}
	var err error
	config.AuthURL, err = url.Parse(authURLStr)
	if err != nil {
		return nil, fmt.Errorf("invalid authorization_url: %w", err)
	}
	if config.ClientID == "" {
		return nil, fmt.Errorf("client_id is required")
	}

	return config, nil
}

func buildAuthURL(config *OIDCProviderConfig, redirectURI, state string) *url.URL {
	params := url.Values{}
	params.Set("client_id", config.ClientID)
	params.Set("response_type", "code")
	params.Set("redirect_uri", redirectURI)
	params.Set("scope", config.Scope)
	params.Set("state", state)

	// Create a shallow copy to avoid mutating the original base URL
	built := *config.AuthURL
	built.RawQuery = params.Encode()
	return &built
}

func generateState(r io.Reader) (string, error) {
	b := make([]byte, 32)
	if _, err := r.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func openBrowser(u *url.URL) error {
	var cmd *exec.Cmd
	urlStr := u.String()
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", urlStr)
	case "linux":
		cmd = exec.Command("xdg-open", urlStr)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	return cmd.Start()
}
