/*
   package database
   db.go
   - contain database pool connection preparation and validation for the established pool connection
*/
package database

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/reshimahendra/lbw-go/internal/config"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/logger"
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

    // prepare context
    ctx := context.Background()

    // tried establish pool connection to database
    pool, err := pgxpool.Connect(ctx, dsn.DSN())
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
