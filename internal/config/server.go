/*
   package config
   config_server.go
   - main configuration for the server
*/
package config

import (
	"strings"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
)

// Server is configuration setup for server
type Server struct {
    // DomainName is domain name for server, ex: mywebsite.com
    DomainName                 string

    // Port is the port used by the server
    Port                       string

    // SecureKey is secret key to make auth token
    SecureKey                  string

    // MinimumSecureKeyLength is the minimum length required for the secret key
    MinimumSecureKeyLength     int

    // AccessTokenExpireDuration is valid duration for the token before expired
    AccessTokenExpireDuration  int64

    // RefreshTokenExpireDuration is a refresh token to request new access token after it expired
    RefreshTokenExpireDuration int64

    // LimitCountPerRequest is the limit count that allowed per request
    LimitCountPerRequest       float64

    // TrustedProxies is all trusted proxies
    TrustedProxies             []string

    // ServerMode is server mode option, value is "production" aor "development"
    ServerMode                 string

    // WelcomeMessage is whether to show the welcome/ greeting when the server is executed
    WelcomeMessage             bool
}

// SetMode is to set server mode 
func (s *Server) SetMode(mode string) (error) {
    mode = strings.ToLower(mode)
    switch mode {
    case "production","development":
        // set ServerMode to new mode
        if s.ServerMode != mode {
            s.ServerMode = mode
        }
    default: 
        return E.New(E.ErrServerMode)
    }

    return nil
}

// GetMode is to get the ServerMode value
func (s *Server) GetMode() (string, error) {
    switch s.ServerMode {
    case "production","development":
        return s.ServerMode, nil    
    }

    return "", E.New(E.ErrServerMode)
}

// GetSecureKey is to get secure key setting
func (s *Server) GetSecureKey() (string, error) {
    if len(s.SecureKey) < s.MinimumSecureKeyLength {
        return "", E.New(E.ErrSecureKeyTooShort)
    }

    return s.SecureKey, nil 
}
