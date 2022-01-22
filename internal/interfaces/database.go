package interfaces

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
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

