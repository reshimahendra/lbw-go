package domain

import "time"

// AuthLoginDTO is 'DTO' (Data Transfer Object) to verify user on login
type AuthLoginDTO struct {
    Email    string `json:"email"`
    Passkey  string `json:"passkey"`
}

// AuthLoginResponse is 'DTO' (Data Transfer Object) to 'Response'
// or sending data to user upon 'login' or request 'refresh token'
type AuthLoginResponse struct {
    AccessToken     string  `json:"access_token"`
    RefreshToken    string  `json:"refresh_token"`
    TransmissionKey string  `json:"transmission_key"`
}

// TokenDetailsDTO is 'DTO' (data Transfer Object) containing
// details of token expiration time
type TokenDetailsDTO struct {
    AccessToken     string  `json:"access_token"`
    RefreshToken    string  `json:"refresh_token"`
    AtExpiresTime   time.Time
    RtExpiresTime   time.Time
    TransmissionKey string  `json:"transmission_key"`
}

