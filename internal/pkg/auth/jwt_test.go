package auth

import (
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/reshimahendra/lbw-go/internal/config"
    E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var wantErr bool = false

type mockJwtToken struct {
    token *jwt.Token
}

func NewMockJwtToken(token *jwt.Token) *mockJwtToken{
    return &mockJwtToken{token}
}

func (m *mockJwtToken) SignedString(key interface{}) (string, error){
    if wantErr {
        return "", E.New(E.ErrTokenInvalid)
    }
    return m.SignedString(key)
}

var (
    aTok = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImV4cCI6MTY0Mzc4OTc5MSwidXNlcl91dWlkIjoidXNlcl91dWlkIiwidXVpZCI6IiJ9.YjxTtQJXnwtHZksr48X67j2eOlkaNNTbtxTH1mDriRE"
    rTok = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImV4cCI6MTY0Mzk2MjU5MSwidXNlcl91dWlkIjoidXNlcl91dWlkIiwidXVpZCI6IiJ9.BgtXywYibwUfP6FZnVXB_e1yJ5KLB6-D8n0DT9V0HRs"
)

// TestCreateToken will testing token create routine
func TestCreateToken(t *testing.T) {
    cases := []struct{
        name, email string
        wantErr bool
    }{
        {"EXPECT SUCCESS", "test@gmail.com", false},
        {"EXPECT FAIL email invalid", "", true},
        {"EXPECT FAIL secure key fail", "test@gmail.com", true},
    }

    err := config.Setup()
    if err != nil {
        t.Errorf("unexpected error occur: %v\n", err)
    }

    _ = config.Get()
    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T){
            if tt.wantErr {
                // prepare mock
                genSecureKey := generateSecureKeyFunc
                generateSecureKeyFunc = func(length int) (string, error) {
                    return "", E.New(E.ErrSecureKeyTooShort)
                }
                defer func() { generateSecureKeyFunc = genSecureKey }()

                // actual test
                got, err := CreateToken(tt.email)

                assert.Error(t, err)
                assert.Nil(t, got)
            } else {
                // actual test
                got, err := CreateToken(tt.email)

                assert.NoError(t, err)
                assert.NotNil(t, got)
            }
        })
    }
}

func TestVerifyToken(t *testing.T) {
    cases := []struct{
        name,token string
        wantErr bool
    }{
        {"EXPECT SUCCESS", aTok, false},
        {"EXPECT FAIL", "fail token", true},
        {"EXPECT FAIL II", "", true},
    }

    err := config.Setup()
    if err != nil {
        t.Fatalf("unexpected error occur: %v\n", err)
    }

    _ = config.Get()
    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T){
            // actual test
            got, err := verifyToken(tt.token)
            if tt.wantErr {
                assert.Error(t, err)
                assert.Nil(t, got)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, got)
            }
        })
    }
}

func TestTokenValid(t *testing.T) {
    cases := []struct{
        name,token string
        wantErr bool
    }{
        {"EXPECT SUCCESS", aTok, false},
        {"EXPECT FAIL invalid token", aTok+"a", true},
        {"EXPECT FAIL", "'", true},
    }

    _ = config.Get()
    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T){
            // actual test
            got, err := TokenValid(tt.token)
            if tt.wantErr {
                if len(tt.token) > 10 {
                    assert.Error(t, err)
                    assert.NotNil(t, got)
                } else {
                    assert.Error(t, err)
                    assert.Nil(t, got)
                }
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, got)
            }
        })
    }
}
