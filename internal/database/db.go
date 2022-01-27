/*
   package database
   db.go
   - contain database pool connection preparation and validation for the established pool connection
*/
package database

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/reshimahendra/lbw-go/internal/config"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// IDatabase is interface to pgxpool method
type IDatabase interface {
	// Exec acquires a connection from the Pool and executes the given SQL.
    // SQL can be either a prepared statement name or an SQL string.
    Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)

    // QueryRow acquires a connection and executes a query that is expected
    // to return at most one row (pgx.Row)
	QueryRow(context.Context, string, ...interface{}) pgx.Row

    // Query acquires a connection and executes a query that returns pgx.Rows.
    // Arguments should be referenced positionally from the SQL string as $1, $2, etc.
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	
    // Close closes all connections in the pool and rejects future Acquire calls. Blocks until all connections are returned
    // to pool and closed.
    Close()    
}

// NewDBPool is to create new pool connection to database
func NewDBPool(dsn config.Database) (*pgxpool.Pool, func(), error) {
    f := func() {}

    // if dsn is invalid, exit process
    if !dsn.IsValid() {
        err := E.New(E.ErrDatabaseConfiguration)
        return nil, f, err 
    }

    // parse connection config
    cfg, err := pgxpool.ParseConfig(dsn.DSN())
    if err != nil {
        logger.Errorf("unable to parse database config: %v", err)
        return nil, f, err
    }
    
    // prepare logrus logger
    logDB := &logrus.Logger{
        Out:          getWriter(),
        Formatter:    &formatter{},
        Hooks:        make(logrus.LevelHooks),
        Level:        logrus.InfoLevel,
        ExitFunc:     os.Exit,
        ReportCaller: false,
    }

    // set logger adapter to logrus
    cfg.ConnConfig.Logger = logrusadapter.NewLogger(logDB)

    // prepare context
    ctx := context.Background()

    // tried establish pool connection to database
    pool, err := pgxpool.ConnectConfig(ctx, cfg)
    if err != nil {
        return nil, f, E.New(E.ErrDatabase)
    }

    // send info log to the logger
    logger.Infof("database connection open on %s:%s", dsn.Hostname, dsn.Port)

    // validate connection. if error occur, return nil
    if err := validateDBPool(ctx, pool); err != nil {
        return nil, f, err
    }

    // return pool instance and other returned values
    return pool, func() {pool.Close()}, nil
}

// validateDBPool is to validate the connection that tried to created by pinging too it
func validateDBPool(ctx context.Context, pool *pgxpool.Pool) error{
    // send ping to make sure the connection is ok
    err := pool.Ping(ctx)
    if err != nil {
        return E.New(E.ErrDatabase)
    }

    // send info to logger notify that the connection established
    logger.Infof("connect to database successfully")
    var (
		currentUser     string
		dbVersion       string
	)

    // get connected user and database version information
	sqlStatement := `select current_user, version();`
	row := pool.QueryRow(ctx, sqlStatement)
	err = row.Scan(&currentUser, &dbVersion)

    // make sure no error occur
	if err != nil {
		return E.New(E.ErrDatabase)
	}

    // send info log to the logger
    logger.Infof("current database user: %s", currentUser)
    logger.Infof("current database: %s", dbVersion)

	return nil
}

// Log writer
func getWriter() io.Writer {
    logDir := "log"
    logPath := filepath.Join(logDir, viper.GetString("logger.databaseLogName"))
    file, err := os.OpenFile(
        logPath,
        os.O_CREATE | os.O_WRONLY | os.O_APPEND, 
        0666)
    if err != nil {
        return os.Stdout
    } else {
        return file
    }
}
// Formatter implements logrus.Formatter interface.
type formatter struct {
	prefix string
}

// Format building log message.
func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var sb bytes.Buffer

	var newLine = "\n"
	if runtime.GOOS == "windows" {
		newLine = "\r\n"
	}

    var s string = ""
    count := 1
    for key, val := range entry.Data {
        s = s + fmt.Sprintf(" %v: %v ", key, val)
        count += 1
        if len(entry.Data) >= count {
            s = s + "|"
        }
    }

	sb.WriteString(strings.ToUpper(entry.Level.String()))
	sb.WriteString(" ")
	sb.WriteString(entry.Time.Format(time.RFC3339))
	sb.WriteString(" ")
	sb.WriteString(f.prefix)
	sb.WriteString(entry.Message)
    sb.WriteString(" ")
    sb.WriteString(s)
	sb.WriteString(newLine)

	return sb.Bytes(), nil
}
