package config

import (
	"testing"

	"github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
    dbname, username, password, hostname, dbport string = "lbw-go", 
        "lotus", "secret", "localhost", "5432"
    sslmode, logmode bool = false, true
    domain, port, secureKey string = "lotusbw.com", "8000", "secret"
    minSecureKeyLen int = 16
    atDur, rtDur int64 = 1, 1
    limitCountReq float64 = 1
    serverMode string = "development"
    welcomeMsg bool = true
    minPassLen int = 8
    dbLog, svrLog, accLog string = ".database.log", ".server.log", ".access.log"

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

func TestGet(t *testing.T) {
    cfg := Get()
    assert.NotNil(t, cfg)

    // make sure secret key and db password are the same first
    // we will assert only equal value here
    secureKey = cfg.Server.SecureKey
    password = cfg.Database.Password

    wantDB := Database{
        DBName   : dbname,
        Username : username,
        Password : password,
        Hostname : hostname,
        Port     : dbport,
        SSLMode  : sslmode,
        LogMode  : logmode,
    }

    wantServer := Server{
        DomainName                 : domain,
        Port                       : port,
        SecureKey                  : secureKey,
        MinimumSecureKeyLength     : minSecureKeyLen,
        AccessTokenExpireDuration  : atDur,
        RefreshTokenExpireDuration : rtDur,
        LimitCountPerRequest       : limitCountReq,
        ServerMode                 : serverMode,
        WelcomeMessage             : welcomeMsg,

    }

    wantAccount := Account{
        MinimumPasswordLength : minPassLen,
    }

    wantLog := Logger{
        DatabaseLogName : dbLog,
        ServerLogName   : svrLog,
        AccessLogName   : accLog,
    }

    assert.Equal(t, wantDB, cfg.Database)
    assert.Equal(t, wantServer, cfg.Server)
    assert.Equal(t, wantAccount, cfg.Account)
    assert.Equal(t, wantLog, cfg.Logger)
}
