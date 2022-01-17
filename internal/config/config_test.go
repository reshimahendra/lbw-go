package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
    dbname, username, password, hostname, dbport string = "booking", 
        "lotus", "secret", "localhost", "5432"
    sslmode, logmode bool = false, true
    domain, port, secretKey string = "lotusbw.com", "8000", "secret"
    minSecureKeyLen int = 16
    atDur, rtDur int64 = 1, 1
    limitCountReq float64 = 1
    serverMode string = "development"
    welcomeMsg bool = true
    minPassLen int = 8
    dbLog, svrLog, accLog string = ".database.log", ".server.log", ".access.log"
)

// TestSetup will test the setup configuration function
func TestSetup(t *testing.T) {
    // prepare to load setting
    err := Setup()
    if err != nil {
        t.Fatalf("unexpected error occur: %v", err)
    }
}

func TestGet(t *testing.T) {
    cfg := Get()
    assert.NotNil(t, cfg)

    // make sure secret key and db password are the same first
    // we will assert only equal value here
    secretKey = cfg.Server.SecretKey
    password = cfg.Database.Password

    wantDB := DatabaseConfiguration{
        DBName   : dbname,
        Username : username,
        Password : password,
        Hostname : hostname,
        Port     : dbport,
        SSLMode  : sslmode,
        LogMode  : logmode,
    }

    wantServer := ServerConfiguration{
        DomainName                 : domain,
        Port                       : port,
        SecretKey                  : secretKey,
        MinimumSecureKeyLength     : minSecureKeyLen,
        AccessTokenExpireDuration  : atDur,
        RefreshTokenExpireDuration : rtDur,
        LimitCountPerRequest       : limitCountReq,
        ServerMode                 : serverMode,
        WelcomeMessage             : welcomeMsg,

    }

    wantAccount := AccountConfiguration{
        MinimumPasswordLength : minPassLen,
    }

    wantLog := LoggerConfiguration{
        DatabaseLogName : dbLog,
        ServerLogName   : svrLog,
        AccessLogName   : accLog,
    }

    assert.Equal(t, cfg.Database, wantDB)
    assert.Equal(t, cfg.Server, wantServer)
    assert.Equal(t, cfg.Account, wantAccount)
    assert.Equal(t, cfg.Logger, wantLog)
}
