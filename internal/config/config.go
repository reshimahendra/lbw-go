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

// Configuration is main configuration wrapper
type Configuration struct{
    // Server is server configuration
    Server Server

    // Database is database configuration
    Database Database
   
    // Account is account configuration
    Account Account

    // Logger is logger configuration
    Logger Logger
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

    // try ro read config file
    if err = viper.ReadInConfig(); err != nil {
        logger.Errorf("error reading config file: %v\n", err)
        return err
    }

    // unmarshal from the yaml file to Configuration struct
    if err = viper.Unmarshal(&c); err != nil {
        logger.Errorf("error unable to decode config into struct: %v\n", err)
        return err
    }

    config = c

    return
}
