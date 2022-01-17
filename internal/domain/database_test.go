/*
   package domain
   - database test
*/
package domain

import (
	"testing"

	"github.com/reshimahendra/lbw-go/internal/config"
	"github.com/stretchr/testify/assert"
)

// TestSetup will test setup configuration of the database
func TestSetup(t *testing.T) {
    // we must load the main setting before loading the database setup
    err := config.Setup()
    if err != nil {
        t.Fatalf("unexpected error occur: %v\n", err)
    }
     
    // loading database setup
    err = Setup()
    if err != nil {
        t.Fatalf("unexpected error occur: %v\n", err)
    }
}

// TestDB will test the get DB function of the database
func TestDB(t *testing.T) {
    conn := DB()

    assert.NotNil(t, conn)
}
