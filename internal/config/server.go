/*
   package config
   config_server.go
   - main configuration for the server
*/
package config

import (
	"errors"
	"strings"
)

// Server is configuration setup for server
type Server struct {
    // DomainName is domain name for server, ex: mywebsite.com
    DomainName                 string

    // Port is the port used by the server
    Port                       string

    // SecretKey is secret key to make auth token
    SecretKey                  string

    // MinimumSecretKeyLength is the minimum length required for the secret key
    MinimumSecretKeyLength     int

    // AccessTokenExpireDuration is valid duration for the token before expired
    AccessTokenExpireDuration  int64

    // RefreshTokenExpireDuration is a refresh token to request new access token after it expired
    RefreshTokenExpireDuration int64

    // LimitCountPerRequest is the limit count that allowed per request
    LimitCountPerRequest       float64

    // ServerMode is server mode option, value is "production" aor "development"
    ServerMode                 string

    // WelcomeMessage is whether to show the welcome/ greeting when the server is executed
    WelcomeMessage             bool
}

// SetMode is to set server mode 
func (s *Server) SetMode(mode string) (error) {
    // mode = strings.ToLower(mode)
    if mode != "production" || mode != "development" {
        return errors.New("server mode value must be 'production' or 'development'")
    }

    // set ServerMode to new mode
    if s.ServerMode != mode {
        s.ServerMode = mode
    }

    return nil
}

// GetMode is to get the ServerMode value
func (s *Server) GetMode() (string, error) {
    if s.ServerMode != "production" || s.ServerMode != "development" {
        return "", errors.New("server mode value must be 'production' or 'development'")
    }

    return s.ServerMode, nil    
}

// GetSecretKey is to get secret key setting
func (s *Server) GetSecretKey() (string, error) {
    if len(s.SecretKey) < s.MinimumSecretKeyLength {
        return "", errors.New("secret key length is less than minimum requirement length")
    }

    return s.SecretKey, nil 
}
