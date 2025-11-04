package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	callbackPath                = "/callback"
	redirectScheme              = "http"
	redirectHost                = "localhost"
	localBindIP                 = "127.0.0.1"
	readHeaderTimeout           = 5 * time.Second
	closeAfterWriteDelay        = 200 * time.Millisecond
	gracefulShutdownTimeout     = 2 * time.Second
	defaultCallbackWaitTimeout  = 5 * time.Minute
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

	ln, err := net.Listen("tcp", net.JoinHostPort(localBindIP, "0"))
	if err != nil {
		return nil, fmt.Errorf("listen on callback port: %w", err)
	}
	port := ln.Addr().(*net.TCPAddr).Port

	redirectURL := &url.URL{
		Scheme: redirectScheme,
		Host:   net.JoinHostPort(redirectHost, fmt.Sprintf("%d", port)),
		Path:   callbackPath,
	}
	redirectURI := redirectURL.String()

	state, err := generateState()
	if err != nil {
		return nil, fmt.Errorf("generate state: %w", err)
	}

    authURL := buildAuthURL(config, redirectURI, state)

    if err := openBrowser(authURL); err != nil {
		log.Warnf("Failed to open browser automatically: %v", err)
	}

    log.Info("\nPlease visit the following URL to authenticate:")
    log.Info(authURL.String())

	authCode, err := waitForCallback(ln, state, defaultCallbackWaitTimeout)
	if err != nil {
		return nil, err
	}

	return structpb.NewStruct(map[string]interface{}{
		"authorization_code": authCode,
		"redirect_uri":       redirectURI,
	})
}

func waitForCallback(ln net.Listener, expectedState string, timeout time.Duration) (string, error) {
    defer ln.Close()

	var code string

		srv := &http.Server{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != callbackPath {
				http.NotFound(w, r)
				return
			}
			q := r.URL.Query()

			if e := q.Get("error"); e != "" {
				sendErrorPage(w, "Authentication Failed", q.Get("error_description"))
					go func() { time.Sleep(closeAfterWriteDelay); _ = ln.Close() }()
				return
			}

			if q.Get("state") != expectedState {
				sendErrorPage(w, "Security Error", "Invalid state parameter")
					go func() { time.Sleep(closeAfterWriteDelay); _ = ln.Close() }()
				return
			}

			c := q.Get("code")
			if c == "" {
				sendErrorPage(w, "Authentication Failed", "No authorization code received")
					go func() { time.Sleep(closeAfterWriteDelay); _ = ln.Close() }()
				return
			}

			code = c
			sendSuccessPage(w)
				go func() { time.Sleep(closeAfterWriteDelay); _ = ln.Close() }()
		}),
			ReadHeaderTimeout: readHeaderTimeout,
	}

	t := time.AfterFunc(timeout, func() { _ = ln.Close() })
	defer t.Stop()

	_ = srv.Serve(ln)

	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
	_ = srv.Shutdown(ctx)
	cancel()

	if code == "" {
		return "", fmt.Errorf("authentication failed or timed out after %s", timeout)
	}
	return code, nil
}

func sendSuccessPage(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, successPageTemplate)
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

    getString := func(m map[string]*structpb.Value, key string) string {
        if v, ok := m[key]; ok && v != nil {
            return v.GetStringValue()
        }
        return ""
    }

    config := &OIDCProviderConfig{
        Name:     getString(fields, "name"),
        ClientID: getString(fields, "client_id"),
        Scope:    getString(fields, "scope"),
    }

    authURLStr := getString(fields, "authorization_url")
    if authURLStr == "" {
        return nil, fmt.Errorf("authorization_url is required")
    }
    u, err := url.Parse(authURLStr)
    if err != nil {
        return nil, fmt.Errorf("invalid authorization_url: %w", err)
    }
    config.AuthURL = u
	if config.ClientID == "" {
		return nil, fmt.Errorf("client_id is required")
	}
	if config.Scope == "" {
		config.Scope = "email"
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

func getFreePort() (int, error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

func generateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
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
