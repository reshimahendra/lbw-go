/*
   package config
   database_test.go
   - test unit for database
*/
package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDatabaseConfig is for testing database config behaviour
func TestDatabaseConfig(t *testing.T) {
    // set ssl enable so the test can enter the if statement at sslmode condition
    wantDB.SSLMode= true

    wantDSN := "postgres://lotus:secret@localhost:5432/lbw-go?sslmode=enable" 
    assert.Equal(t, wantDSN, wantDB.DSN())
    assert.Equal(t, true, wantDB.IsValid())
}
