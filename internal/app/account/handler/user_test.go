/*
   package handler
   user_test.go
   - testing behaviour of user handler
*/
package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/reshimahendra/lbw-go/internal/config"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/pkg/auth"
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
            StatusID  : 0,
            RoleID    : 0,
            CreatedAt : time.Now(),
            UpdatedAt : time.Now(),
        },
    }

    wantErrInactive bool

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

// GetByEmail is mocked GetByEmail method to satisfy IUserService interface
func (m *mockUserHandler) GetByEmail(email string) (*d.UserCredential, error) {
    if wantErr {
        return nil, E.New(E.ErrDataNotFound)
    }
    if email=="reshi@lotusbw.com" {
        return &d.UserCredential{
            ID : u[0].ID,
            Username : "reshi",
            PassKey : "$2a$14$t5Bf3SLtsyazg2nzQ57HyeDMLsHGvm2x/VyjmM5XGojiPj4WmWDhi", 
            StatusID : 1,
        }, nil
    } else if email=="inactive@lotusbw.com" {
        return &d.UserCredential{
            ID : u[0].ID,
            Username : "reshi",
            PassKey : "$2a$14$t5Bf3SLtsyazg2nzQ57HyeDMLsHGvm2x/VyjmM5XGojiPj4WmWDhi", 
            StatusID : 0,
        }, nil
    }   
    return &d.UserCredential{
        ID : u[0].ID,
        Username : u[0].Username,
        PassKey : "$2a$14$t5Bf3SLtsyazg2nzQ57HyeDMLsHGvm2x/VyjmM5XGojiPj4WmWDhi", 
        StatusID : u[0].StatusID,
    }, nil
}

// GetCredential is mocked GetCredential method to satisfy IUserService interface
func (m *mockUserHandler) GetCredential(username,passkey string) (*d.UserCredential, error) {
    if wantErr {
        return nil, E.New(E.ErrDataNotFound)
    }
    return &d.UserCredential{
        ID : u[0].ID,
        Username : u[0].Username,
        PassKey : "secret", 
        StatusID : u[0].StatusID,
    }, nil
}

// IsUserExist is mocked IsUserExist method to satisfy IUserService interface
func (m *mockUserHandler) IsUserExist(username,email string) bool {
    return wantErr
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

// TestSignupHandler will test behaviour of Signup method of handler layer
func TestSignupHandler(t *testing.T) {
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
        handler.SignupHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(u[0])
        assert.NoError(t, err)

        // validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]))
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success signup") 
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
        handler.SignupHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusBadRequest, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrRequestDataInvalidMsg)
    })

    // EXPECT FAIL user exist error. Simulation done by forcing IsUserExist resulting error
    // by setting wantErr=true
    t.Run("EXPECT FAIL user exist error", func(t *testing.T){
        // prepare request/ response / gin context
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

        // set content type to json.
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call. Force IsUserExist resulting true to generate error
        wantErr=true
        handler.SignupHandler(context)
        wantErr=false

        // validation and verification
        assert.Equal(t, http.StatusBadRequest, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrUserAlreadyRegisteredMsg)
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
        handler.SignupHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusInternalServerError, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrDataIsInvalidMsg)
    })
}

// TestSigninHandler will test behaviour of Signin method of handler layer
func TestSigninHandler(t *testing.T) {
    // prepare the test handler 
    handler := NewTestUserHandler(t)
    
    // prepare config
    err := config.Setup()
    assert.NoError(t, err)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // prepare mock with ur[0] values
        req := new(d.AuthLoginDTO)
        req.Email  = "reshi@lotusbw.com" 
        req.Passkey= "12345678"

        uJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.SigninHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success signin")
    })

    // EXPECT FAIL bind json error. Simulation done by removing request body so 
    // reading request data will be error
    t.Run("EXPECT FAIL bind json error", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()

        // inject json to request body
        var err error
        context.Request, err = http.NewRequest("POST", "/", nil)
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.SigninHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusUnauthorized, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrRequestDataInvalidMsg)
    })

    // EXPECT FAIL invalid email error. Simulation done by giving invalid email input 
    t.Run("EXPECT FAIL invalid email error", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()

        // prepare mock with ur[0] values
        req := new(d.AuthLoginDTO)
        req.Email  = "dddd.com"
        req.Passkey= "12345678"

        uJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.SigninHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusInternalServerError, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrTokenCreateMsg)
    })

    // EXPECT FAIL post data error. Simulation done by feeding invalid account data
    t.Run("EXPECT FAIL insert record error", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()

        // prepare mock with ur[0] values
        req := new(d.AuthLoginDTO)
        req.Email  = "dddd.com"
        req.Passkey= "12345678"

        uJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        wantErr=true
        handler.SigninHandler(context)
        wantErr=false

        // validation and verification
        assert.Equal(t, http.StatusUnauthorized, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrDataNotFoundMsg)
    })

    // EXPECT FAIL inactive account error. Simulation done by feeding an inactive account
    t.Run("EXPECT FAIL inactive account error", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // prepare mock with ur[0] values
        req := new(d.AuthLoginDTO)
        req.Email  = "inactive@lotusbw.com" 
        req.Passkey= "12345678"

        uJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.SigninHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusUnauthorized, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrUserNotActiveMsg)
    })

    // EXPECT FAIL password not match error. Simulation done by mocking 
    // CheckPasswordHash function on helper package
    t.Run("EXPECT FAIL password not match error", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // prepare mock with ur[0] values
        req := new(d.AuthLoginDTO)
        req.Email  = "reshi@lotusbw.com" 
        req.Passkey= "12345678"

        uJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // mock helper.CheckPasswordHash
        checkPasswordHash := checkPasswordHashFunc
        checkPasswordHashFunc = func(password string, hash string) bool {
            return false
        }
        defer func() { checkPasswordHashFunc = checkPasswordHash }()

        // actual method handler call
        handler.SigninHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusUnauthorized, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrSignInMsg)
    })
}

// TestRefreshTokenHandler will test behaviour of RefreshTokenHandler method of handler layer
func TestRefreshTokenHandler(t *testing.T) {
    // prepare the test handler 
    handler := NewTestUserHandler(t)
    
    // prepare config
    err := config.Setup()
    assert.NoError(t, err)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // prepare mock token
        testToken, err := auth.CreateToken("test@token.com")
        assert.NoError(t, err)

        uJSON, err := json.Marshal(testToken)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.RefreshTokenHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success refresh token")
    })

    // EXPECT FAIL decode json error. simulated by feeding 'nil' data
    t.Run("EXPECT FAIL decode json error", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer([]byte("invalid"))) 
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.RefreshTokenHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusUnprocessableEntity, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrTokenRefreshMsg)
    })

    // EXPECT FAIL token invalid. simulated by feeding invalid json data
    t.Run("EXPECT FAIL token invalid error", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // prepare mock data with invalid value
        type trashData struct{
            hello string
            datainvalid string
        }
        invalidData := trashData{hello: "invalid", datainvalid:"data"}

        uJSON, err := json.Marshal(invalidData)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.RefreshTokenHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusUnauthorized, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrTokenInvalidMsg)
    })

    // EXPECT FAIL create token error. simulated by feeding invalid email 
    t.Run("EXPECT FAIL create token error", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // prepare mock token
        testToken, err := auth.CreateToken("test@token.com")
        assert.NoError(t, err)

        uJSON, err := json.Marshal(testToken)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // mock auth.CreateToken function to return error
        createToken := createTokenFunc
        createTokenFunc = func(email string) (*d.TokenDetailsDTO, error) {
            return nil, E.New(E.ErrTokenCreate)
        }
        defer func() { createTokenFunc = createToken }()

        // actual method handler call
        handler.RefreshTokenHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusInternalServerError, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrTokenCreateMsg)
    })
}

// TestCheckTokenHandler will test behaviour of CheckTokenHandler method of handler layer
func TestCheckTokenHandler(t *testing.T) {
    // prepare the test handler 
    handler := NewTestUserHandler(t)
    
    // prepare config
    err := config.Setup()
    assert.NoError(t, err)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", nil) 
        assert.NoError(t, err)

        // prepare mock token
        testToken, err := auth.CreateToken("test@token.com")
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")
        context.Request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", testToken.AccessToken))

        // actual method handler call
        handler.CheckTokenHandler(context)
    
        // validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success checking token")
    })

    // EXPECT FAIL token not found error. Simulated by removing 'Authorization' header
    t.Run("EXPECT FAIL token not found", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", nil) 
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.CheckTokenHandler(context)
    
        // validation and verification
        assert.Equal(t, http.StatusUnauthorized, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrTokenNotFoundMsg)
    })

    // EXPECT FAIL token invalid error. Simulated by giving invalid token value 
    t.Run("EXPECT FAIL token invalid error", func(t *testing.T){
        // prepare request/ response / gin context
        // this is shared func, it is located in user.role_test.go
        writer, context := NewTestWriterContext()

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", nil) 
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")
        context.Request.Header.Add("Authorization", "Bearer ahahahahaha_invalid_token_akahahahaha")

        // actual method handler call
        handler.CheckTokenHandler(context)
    
        // validation and verification
        assert.Equal(t, http.StatusUnauthorized, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrTokenInvalidMsg)
    })
}
