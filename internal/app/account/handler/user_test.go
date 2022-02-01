/*
   package handler
   user_test.go
   - testing behaviour of user handler
*/
package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var(
    // u is user mock data slice for User
    u = []*d.UserResponse{
        {
            ID        : uuid.New(),
            Username  : "leonard",
            Firstname : "Leo",
            Lastname  : "Singa",
            Email     : "leo@gmail.com",
            StatusID  : 1,
            RoleID    : 1,
            CreatedAt : time.Now(),
            UpdatedAt : time.Now(),
        },
        {
            ID        : uuid.New(),
            Username  : "jennydoe",
            Firstname : "Jenny",
            Lastname  : "Doe",
            Email     : "jenny@gmail.com",
            StatusID  : 1,
            RoleID    : 0,
            CreatedAt : time.Now(),
            UpdatedAt : time.Now(),
        },
    }

)

// mockUserHandler is mocked user handler for our user service interface
type mockUserHandler struct {
    t *testing.T
}

// NewMockUserHandler is new instance to our mockUserHandler
func NewMockUserHandler(t *testing.T) *mockUserHandler{
    return &mockUserHandler{t}
}

// Create is mocked Create method of IUserService.Create
func (m *mockUserHandler) Create(input d.UserRequest) (*d.UserResponse, error) {
    // simulate input invalid
    if !input.IsValid() {
        return nil, E.New(E.ErrDataIsInvalid)
    }
    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    return u[0], nil
}

// Get is mocked Get method of IUserService.Get
func (m *mockUserHandler) Get(id string) (*d.UserResponse, error) {
    // return nil if force error set to true
    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    return u[1], nil
}

// Gets is mocked Gets method of IUserService.Gets
func (m *mockUserHandler) Gets() ([]*d.UserResponse, error) {
    // return nil if force error set to true
    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    return u, nil
}

// Update is mocked Update method of IUserService.Update
func (m *mockUserHandler) Update(id string, input d.UserRequest) (*d.UserResponse, error) {
    // return nil when input invalid
    if !input.IsValid() {
        return nil, E.New(E.ErrRequestDataInvalid)
    }

    // return nil if force error set to true
    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    return u[0], nil
}

// Delete is mocked Delete method of IUserService.Delete
func (m *mockUserHandler) Delete(id string) (*d.UserResponse, error) {
    // return nil if force error set to true
    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    return u[0], nil
}

// NewTestUserHandler is function wrapper to get the mock handler of our handler layer
func NewTestUserHandler(t *testing.T) *UserHandler{
    t.Helper()

    // set gin to test mode
    gin.SetMode(gin.TestMode)

    // prepare mock
    mock := NewMockUserHandler(t)
    handler := NewUserHandler(mock)

    // return mocked handler
    return handler
}

// TestUserCreateHandler will test behaviour of UserCreateHandler method of handler layer
func TestUserCreateHandler(t *testing.T) {
    // prepare the test handler 
    handler := NewTestUserHandler(t)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // prepare mock with ur[0] values
        req := new(d.UserRequest)
        req.Username  = u[0].Username
        req.Firstname = u[0].Firstname
        req.Lastname  = u[0].Lastname
        req.Email     = u[0].Email
        req.PassKey   = "secret"
        req.StatusID  = u[0].StatusID
        req.RoleID    = u[0].RoleID

        uJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.UserCreateHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(u[0])
        assert.NoError(t, err)

        // validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]))
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success inserting user data") 
    })

    // EXPECT FAIL bind json error. Simulation done by removing request body so 
    // reading request data will be error
    t.Run("EXPECT FAIL bind json error", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()

        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("POST", "/", nil)
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.UserCreateHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusBadRequest, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrRequestDataInvalidMsg)
    })

    // EXPECT FAIL insert error. Simulation done by removing some required field to make 
    // request data invalid
    t.Run("EXPECT FAIL insert record error", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()

        // prepare mock with ur[0] values
        req := new(d.UserRequest)
        req.Firstname = u[0].Firstname
        req.Lastname  = u[0].Lastname
        req.Email     = u[0].Email
        req.PassKey   = "secret"

        uRoleJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uRoleJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.UserCreateHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusInternalServerError, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrDataIsInvalidMsg)

    })
}

// TestUserGetHandler will test behaviour of UserGetHandler
func TestUserGetHandler(t *testing.T) {
    // prepare the test
    handler := NewTestUserHandler(t)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value: u[1].ID.String()},
        }

        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("GET", "/:id", nil)
        assert.NoError(t, err)

        // actual method call
        handler.UserGetHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(u[1])
        assert.NoError(t, err)

        // test validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]))
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success getting user data")
    })

    // EXPECT FAIL get data error. Simulated by inserting non existing id
    t.Run("EXPECT FAIL get data error", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:u[1].ID.String()},
        }

        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("GET", "/:id", nil)
        assert.NoError(t, err)

        // actual method call
        // set wantErr to force error return
        wantErr = true
        handler.UserGetHandler(context)
        wantErr = false

        assert.Equal(t, http.StatusInternalServerError, writer.Code)
    })
}

// TestUserGetsHandler will test behaviour of UserGetsHandler
func TestUserGetsHandler(t *testing.T) {
    // prepare the test
    handler := NewTestUserHandler(t)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("GET", "/", nil)
        assert.NoError(t, err)

        // actual method call
        handler.UserGetsHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(u)
        assert.NoError(t, err)

        // test validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]))
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success getting user data")
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
        handler.UserGetsHandler(context)
        wantErr = false

        assert.Equal(t, http.StatusInternalServerError, writer.Code)
    })
}

// TestUserUpdateHandler will test behaviour of UserUpdateHandler method of handler layer
func TestUserUpdateHandler(t *testing.T) {
    // prepare the test handler 
    handler := NewTestUserHandler(t)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:u[0].ID.String()},
        }

        // prepare mock with ur[0] values
        req := new(d.UserRequest)
        req.Username  = u[0].Username
        req.Firstname = u[0].Firstname
        req.Lastname  = u[0].Lastname
        req.Email     = u[0].Email
        req.PassKey   = "secret"
        req.StatusID  = u[0].StatusID
        req.RoleID    = u[0].RoleID

        uJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("PUT", "/:id", bytes.NewBuffer(uJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.UserUpdateHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(u[0])
        assert.NoError(t, err)

        // validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]))
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success updating user data") 
    })

    // EXPECT FAIL bind json error. Simulation done by removing request body so 
    // reading request data will be error
    t.Run("EXPECT FAIL bind json error", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()

        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:u[0].ID.String()},
        }

        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("PUT", "/:id", nil)
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.UserUpdateHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusBadRequest, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrRequestDataInvalidMsg)
    })

    // EXPECT FAIL insert error. Simulation done by removing some required field to make 
    // request data invalid
    t.Run("EXPECT FAIL insert record error", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()

        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:u[0].ID.String()},
        }

        // prepare mock with ur[0] values
        req := new(d.UserRequest)
        req.Firstname = u[0].Firstname
        req.Lastname  = u[0].Lastname
        req.Email     = u[0].Email
        req.PassKey   = "secret"

        uRoleJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("PUT", "/:id", bytes.NewBuffer(uRoleJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.UserUpdateHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusInternalServerError, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrRequestDataInvalidMsg)
    })
}

// TestUserDeleteHandler will test behaviour of UserDeleteHandler
func TestUserDeleteHandler(t *testing.T) {
    // prepare the test
    handler := NewTestUserHandler(t)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value: u[0].ID.String()},
        }

        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("DELETE", "/:id", nil)
        assert.NoError(t, err)

        // actual method call
        handler.UserDeleteHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(u[0])
        assert.NoError(t, err)

        // test validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]))
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success deleting user data")
    })

    // EXPECT FAIL get data error. Simulated by inserting non existing id
    t.Run("EXPECT FAIL get data error", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:u[0].ID.String()},
        }

        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("DELETE", "/:id", nil)
        assert.NoError(t, err)

        // actual method call
        // set wantErr to force error return
        wantErr = true
        handler.UserDeleteHandler(context)
        wantErr = false

        assert.Equal(t, http.StatusInternalServerError, writer.Code)
    })
}
