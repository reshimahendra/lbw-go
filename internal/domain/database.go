package domain

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var db *sql.DB

func DB() *sql.DB {
    return db
}

func Setup() (err error) {
    username := viper.GetString("database.username")
    password := viper.GetString("database.password")
    host     := viper.GetString("database.hostname")
    port     := viper.GetString("database.port")
    dbname   := viper.GetString("database.dbname")
    // sslmode  := viper.GetBool("database.ssl_mode")
    // logmode  := viper.GetBool("database.logmode")

    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbname)

    database, err := sql.Open("postgres", dsn)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database:%v\n", err)
        os.Exit(1)
    }

    defer database.Close()

    if err := database.Ping(); err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database:%v\n", err)
        os.Exit(1)
    }

    db = database

    return
}
