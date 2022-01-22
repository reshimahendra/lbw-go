/*
    package datastore
    db.go
    - contain database pool connection preparation and validation for the established pool connection
*/
package datastore

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/reshimahendra/lbw-go/internal/config"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/logger"
)

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
