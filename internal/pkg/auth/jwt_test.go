package auth

import (
	"testing"

	"github.com/reshimahendra/lbw-go/internal/config"
	"github.com/stretchr/testify/assert"
)

// TestCreateToken will testing token create routine
func TestCreateToken(t *testing.T) {
    t.Skip()
    cases := []struct{
        name, email string
        wantErr bool
    }{
        {"SuccessA", "test@gmail.com", false},
        {"FailA", "test.com", true},
    }

    err := config.Setup()
    if err != nil {
        t.Fatalf("unexpected error occur: %v\n", err)
    }

    _ = config.Get()
    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T){
            got, err := CreateToken(tt.email)
            if err != nil {
                t.Fatalf("error unexpected: %v", err)
            }
            t.Logf("GOT: %v", got)
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
