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

type OIDCProviderConfig struct {
	Name     string
	AuthURL  string
	ClientID string
	Scope    string
}

type OIDCProvider struct{}

func (p *OIDCProvider) Authenticate(providerData *structpb.Struct) (*structpb.Struct, error) {
	config, err := parseProviderData(providerData)
	if err != nil {
		return nil, fmt.Errorf("parse provider data: %w", err)
	}

	port, err := getFreePort()
	if err != nil {
		return nil, fmt.Errorf("find free port: %w", err)
	}
	redirectURI := fmt.Sprintf("http://localhost:%d/callback", port)

	state, err := generateState()
	if err != nil {
		return nil, fmt.Errorf("generate state: %w", err)
	}

	authURL := buildAuthURL(config, redirectURI, state)

	if err := openBrowser(authURL); err != nil {
		log.Warnf("Failed to open browser automatically: %v", err)
		log.Info("\nPlease visit the following URL to authenticate:")
		log.Info(authURL)
	}

	authCode, err := waitForCallback(port, state, 5*time.Minute)
	if err != nil {
		return nil, err
	}

	return structpb.NewStruct(map[string]interface{}{
		"authorization_code": authCode,
		"redirect_uri":       redirectURI,
	})
}

func waitForCallback(port int, expectedState string, timeout time.Duration) (string, error) {
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return "", fmt.Errorf("listen: %w", err)
	}
	defer ln.Close()

	var code string

	srv := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/callback" {
				http.NotFound(w, r)
				return
			}
			q := r.URL.Query()

			if e := q.Get("error"); e != "" {
				sendErrorPage(w, "Authentication Failed", q.Get("error_description"))
				go func() { time.Sleep(200 * time.Millisecond); _ = ln.Close() }()
				return
			}

			if q.Get("state") != expectedState {
				sendErrorPage(w, "Security Error", "Invalid state parameter")
				go func() { time.Sleep(200 * time.Millisecond); _ = ln.Close() }()
				return
			}

			c := q.Get("code")
			if c == "" {
				sendErrorPage(w, "Authentication Failed", "No authorization code received")
				go func() { time.Sleep(200 * time.Millisecond); _ = ln.Close() }()
				return
			}

			code = c
			sendSuccessPage(w)
			go func() { time.Sleep(200 * time.Millisecond); _ = ln.Close() }()
		}),
		ReadHeaderTimeout: 5 * time.Second,
	}

	t := time.AfterFunc(timeout, func() { _ = ln.Close() })
	defer t.Stop()

	_ = srv.Serve(ln)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_ = srv.Shutdown(ctx)
	cancel()

	if code == "" {
		return "", fmt.Errorf("authentication failed or timed out after %s", timeout)
	}
	return code, nil
}

func sendSuccessPage(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<!DOCTYPE html><html><head><meta charset="utf-8"><title>Success</title><script>setTimeout(function(){ window.close(); }, 800);</script></head><body style="font-family:Segoe UI,Arial,sans-serif;text-align:center;padding-top:10%;background:#f5f9ff;"><div style="display:inline-block;background:#fff;border-radius:10px;padding:30px 40px;box-shadow:0 4px 12px rgba(0,0,0,0.1);"><div style="font-size:40px;color:#4CAF50;"></div><h2 style="color:#1a73e8;margin:10px 0;">Authentication Successful</h2><p style="color:#555;">You can close this window.</p></div></body></html>`)
}

func sendErrorPage(w http.ResponseWriter, title, message string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, `<!DOCTYPE html><html><head><meta charset="utf-8"><title>%s</title></head><body style="font-family:Segoe UI,Arial,sans-serif;text-align:center;padding-top:10%%;"><h2>✗ %s</h2><p>%s</p></body></html>`, title, title, message)
}

func parseProviderData(data *structpb.Struct) (*OIDCProviderConfig, error) {
	fields := data.GetFields()

	config := &OIDCProviderConfig{
		Name:     fields["name"].GetStringValue(),
		AuthURL:  fields["authorization_url"].GetStringValue(),
		ClientID: fields["client_id"].GetStringValue(),
		Scope:    fields["scope"].GetStringValue(),
	}

	if config.AuthURL == "" {
		return nil, fmt.Errorf("authorization_url is required")
	}
	if config.ClientID == "" {
		return nil, fmt.Errorf("client_id is required")
	}
	if config.Scope == "" {
		config.Scope = "email"
	}

	return config, nil
}

func buildAuthURL(config *OIDCProviderConfig, redirectURI, state string) string {
	params := url.Values{}
	params.Set("client_id", config.ClientID)
	params.Set("response_type", "code")
	params.Set("redirect_uri", redirectURI)
	params.Set("scope", config.Scope)
	params.Set("state", state)

	return fmt.Sprintf("%s?%s", config.AuthURL, params.Encode())
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

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	return cmd.Start()
}
