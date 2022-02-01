/*
   package config
   server_test.go
   - test unit for server
*/
package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestServerConfig is for testing server config behaviour
func TestServerConfig(t *testing.T) {
    // make copy of wantServer
    tmpServer := wantServer

    // test SetMode method with typical/ expected behaviour
    assert.NoError(t, tmpServer.SetMode("production")) 
    mode, err := tmpServer.GetMode()
    assert.NoError(t, err)
    assert.Equal(t, "production", mode)

    // test GetMode method with typical/ normal operation
    assert.NoError(t, tmpServer.SetMode("development")) 
    mode, err = tmpServer.GetMode()
    assert.NoError(t, err)
    assert.Equal(t, "development", mode)

    // test SetMode method with error return
    assert.Error(t, tmpServer.SetMode("noavail")) 

    // test GetMode method with error return
    tmpServer.ServerMode = "nothing"
    mode, err = tmpServer.GetMode()
    assert.Error(t, err)
    assert.Equal(t, "", mode)

    // test GetSecureKey method with typical/ normal operation
    key, err := tmpServer.GetSecureKey()
    assert.NoError(t, err)
    assert.Equal(t, tmpServer.SecureKey, key)

    // test GetSecureKey method with error return
    tmpServer.SecureKey = "invalidkey"
    key, err = tmpServer.GetSecureKey()
    assert.Error(t, err)
    assert.Equal(t, "", key)
}
