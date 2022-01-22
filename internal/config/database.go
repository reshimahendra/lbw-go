/*
   package config
   config_database.go
   - main configuration for postgresql database
*/
package config

import "fmt"

// DatabaseConfiguration is configuration setup for database
type Database struct {
    // DBName is database name
    DBName   string

    // Username is database username
    Username string

    // Password is database password that match with the database user
    Password string

    // Hostname is database hostname
    Hostname string

    // Port is the port that listen to the database connection
    Port     string

    // SSLMode is ssl connection option to database
    SSLMode  bool

    // LogMode is logging option to log the database operation
    LogMode  bool
}

// DSN will get the datasource name of the database connection
func (db *Database) DSN() string {
    // default sslmode set to disable 
    sslmode := "disable" 

    // if ssl mode enabled
    if db.SSLMode {
        sslmode = "enable"
    }

    // return dns string/ connection uri
    return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", 
        db.Username,
        db.Password,
        db.Hostname,
        db.Port,
        db.DBName,
        sslmode,
    )
}

// IsValid is to check whether databse configuration is valid
func (db *Database) IsValid() bool {
    return db.Username != "" &&
        db.Password != "" &&
        db.Hostname != "" &&
        db.Port != "" &&
        db.DBName != ""
}
