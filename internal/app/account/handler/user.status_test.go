/*
   package handler
   user.status_test.go
   - testing behaviour of user.status handler
*/
package handler

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/stretchr/testify/assert"
)
var (
    us = []*d.UserStatus{
        {ID:0, StatusName:"inactive",Description:"inactive member"},
        {ID:1, StatusName:"active",Description:"active member"},
        {ID:2, StatusName:"vip",Description:"vip member"},
    }
)

// mockUserStatusHandler is mocked user.status handler for our user.status service interface
type mockUserStatusHandler struct {
    t *testing.T
}

// NewMockUserStatusHandler is new instance to our mockUserStatusHandler
func NewMockUserStatusHandler(t *testing.T) *mockUserStatusHandler{
    return &mockUserStatusHandler{t}
}

// Get is mocked Get method of IUserStatusService.Get
func (m *mockUserStatusHandler) Get(id int) (*d.UserStatus, error) {
    if len(us) < id {
        return nil, E.New(E.ErrParamIsInvalid)
    }

    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    return us[id], nil
}

// Gets is mocked Gets method of IUserStatusService.Gets
func (m *mockUserStatusHandler) Gets() ([]*d.UserStatus, error) {
    // return nil if force error set to true
    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    return us, nil
}

// NewTestUserStatusHandler is function wrapper to get the mock handler of our handler layer
func NewTestUserStatusHandler(t *testing.T) *UserStatusHandler{
    t.Helper()

    // set gin to test mode
    gin.SetMode(gin.TestMode)

    // prepare mock
    mock := NewMockUserStatusHandler(t)
    handler := NewUserStatusHandler(mock)

    // return mocked handler
    return handler
}

// TestUserStatusHandler will test behaviour of handler available in UserStatusHandler
func TestUserStatusHandler(t *testing.T) {
    // prepare the test
    handler := NewTestUserStatusHandler(t)

    // EXPECT SUCCESS GET will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS GET", func(t *testing.T){
        // prepare request/ response / gin context
        // this function is shared from "user.role_test.go"
        writer, context := NewTestWriterContext()
        
        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:"1"},
        }

        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("GET", "/:id", nil)
        assert.NoError(t, err)

        // actual method call
        handler.UserStatusGetHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(us[1])
        assert.NoError(t, err)

        // test validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]))
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success get user.status data")
    })

    // EXPECT FAIL GET bad param id. Simulated by remove the id param
    t.Run("EXPECT FAIL GET bad param id", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("GET", "/:id", nil)
        assert.NoError(t, err)

        // actual method call
        handler.UserStatusGetHandler(context)

        assert.Equal(t, http.StatusBadRequest, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrParamIsInvalidMsg)
    })

    // EXPECT FAIL GET empty data error. Simulated by demanding non existing id
    t.Run("EXPECT FAIL GET empty data error", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:"1"},
        }

        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("GET", "/:id", nil)
        assert.NoError(t, err)

        // actual method call
        // set wantErr to force error return
        wantErr = true
        handler.UserStatusGetHandler(context)
        wantErr = false

        assert.Equal(t, http.StatusInternalServerError, writer.Code)
    })

    // EXPECT SUCCESS GETS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS GETS", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()

        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("GET", "/", nil)
        assert.NoError(t, err)

        // actual method call
        handler.UserStatusGetsHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(us)
        assert.NoError(t, err)

        // test validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success get user.status data")
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]) )
    })

    // EXPECT FAIL get data error. Simulated by inserting non existing id
    t.Run("EXPECT FAIL get data error", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("GET", "/", nil)
        assert.NoError(t, err)

        // actual method call
        // set wantErr to force error return
        wantErr = true
        handler.UserStatusGetsHandler(context)
        wantErr = false

        assert.Equal(t, http.StatusInternalServerError, writer.Code)
    })
}
