/*
   Package config
   - main configuration routine
*/
package config

import (
	"path/filepath"
	"runtime"

	"github.com/reshimahendra/lbw-go/internal/pkg/logger"
	"github.com/spf13/viper"
)

var (
    // config is local variable which will passed to Get() function
    config *Configuration

    // get the root directory of our project
    _, base, _, _ = runtime.Caller(0)
    basePath = filepath.Join(filepath.Dir(base), "../..")
)

// DatabaseConfiguration is configuration setup for database
type DatabaseConfiguration struct {
    DBName   string
    Username string
    Password string
    Hostname string
    Port     string
    SSLMode  bool
    LogMode  bool
}

// ServerConfiguration is configuration setup for server
type ServerConfiguration struct {
    DomainName                 string
    Port                       string
    SecretKey                  string
    MinimumSecureKeyLength     int
    AccessTokenExpireDuration  int64
    RefreshTokenExpireDuration int64
    LimitCountPerRequest       float64
    ServerMode                 string
    WelcomeMessage             bool
}

// AccountConfiguration is configuration setup for user account
type AccountConfiguration struct {
    MinimumPasswordLength int
}

// LoggerConfiguration is configuration setup for logger/ log generator
type LoggerConfiguration struct {
    DatabaseLogName string
    ServerLogName   string
    AccessLogName   string
}

// Configuration is main configuration wrapper
type Configuration struct{
    Server ServerConfiguration
    Database DatabaseConfiguration
    Account AccountConfiguration
    Logger LoggerConfiguration
}

// Get will get configuration setting
func Get() *Configuration {
    return config
}

// Setup will initiate main configuration
func Setup() (err error) {
    var c *Configuration

    // locate the configuration file
    viper.SetConfigName(".config.yaml")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(basePath)
    viper.AddConfigPath(filepath.Join(basePath, "config"))
    viper.AddConfigPath("./config")

    // try ro read config
    if err = viper.ReadInConfig(); err != nil {
        logger.Errorf("error reading config file: %v\n", err)
        return err
    }

    if err = viper.Unmarshal(&c); err != nil {
        logger.Errorf("error unable to decode config into struct: %v\n", err)
        return err
    }

    config = c

    return
}
