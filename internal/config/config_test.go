/*
   package config
   config_test.go
   - testing behaviour of database config
*/
package config

import (
	"testing"

	"github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
    // wantDB is temporary Database configuration test value
    wantDB = Database{
        DBName   : "lbw-go",
        Username : "lotus",
        Password : "secret",
        Hostname : "localhost",
        Port     : "5432",
        SSLMode  : false,
        LogMode  : true,
    }

    // wantServer is temporary server configuration test value
    wantServer = Server{
        DomainName                 : "lotusbw.com",
        Port                       : "8000",
        SecureKey                  : "secure-key-is-a-secret",
        MinimumSecureKeyLength     : 16,
        AccessTokenExpireDuration  : 1,
        RefreshTokenExpireDuration : 1,
        LimitCountPerRequest       : 1,
        TrustedProxies             : []string{"127.0.0.1","localhost"},
        ServerMode                 : "production",
        WelcomeMessage             : true,

    }

    // wantAccount is temporary account configuration test value
    wantAccount = Account{
        MinimumPasswordLength : 8,
    }

    // wantLog is temporary logger configuration test value
    wantLog = Logger{
        DatabaseLogName : ".database.log",
        ServerLogName   : ".server.log",
        AccessLogName   : ".access.log",
    }


    // mock func
    viperReadInConfigFunc = viperReadInConfig
    viperUnmarshalFunc = viperUnmarshal
)

// TestSetup will test the setup configuration function
func TestSetup(t *testing.T) {
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare to load setting
        err := Setup()
        if err != nil {
            t.Fatalf("unexpected error occur: %v", err)
        }
    })

    t.Run("EXPECT FAIL read config error", func(t *testing.T){
        viperReadInConfig = func() error {
            return errors.New(errors.ErrDataIsInvalid)
        }
        defer func(){
            viperReadInConfig = viperReadInConfigFunc
        }()

        err := Setup()
        assert.Error(t, err)
    })

    t.Run("EXPECT FAIL unmarshal error", func(t *testing.T){
        viperUnmarshal = func(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
            return errors.New(errors.ErrDataIsInvalid)
        }
        defer func(){
            viperUnmarshal = viperUnmarshalFunc
        }()

        err := Setup()
        assert.Error(t, err)
    })
}

// TestGet is test for Get() function behaviour
func TestGet(t *testing.T) {
    cfg := Get()
    assert.NotNil(t, cfg)

    // make sure secret key and db password are the same first
    // we will assert only equal value here
    cfg.Server.SecureKey = wantServer.SecureKey
    cfg.Database.Password = wantDB.Password

    assert.Equal(t, wantDB, cfg.Database)
    assert.Equal(t, wantServer, cfg.Server)
    assert.Equal(t, wantAccount, cfg.Account)
    assert.Equal(t, wantLog, cfg.Logger)
}
