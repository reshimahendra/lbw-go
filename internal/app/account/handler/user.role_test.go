package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/stretchr/testify/assert"
)
var (
    wantErr bool = false
    urHeader = []string{"id","role_name","description","created_at","updated_at"}
    ur = []*d.UserRole{
        {ID : 0, RoleName: "Guest", Description: "Guest role", CreatedAt: time.Now(), UpdatedAt: time.Now()},
        {ID : 1, RoleName: "Superuser", Description: "Superuser role", CreatedAt: time.Now(), UpdatedAt: time.Now()},
        {ID : 2, RoleName: "Admin", Description: "Admin role", CreatedAt: time.Now(), UpdatedAt: time.Now()},
    }
    errUser = d.UserRole{ID : 3, RoleName: "FAIL", Description: "FAIL role", CreatedAt: time.Now()}
)

// mockUserRoleHandler is mocked user.role handler for our user.role service interface
type mockUserRoleHandler struct {
    t *testing.T
}

// NewMockUserRoleHandler is new instance to our mockUserRoleHandler
func NewMockUserRoleHandler(t *testing.T) *mockUserRoleHandler{
    return &mockUserRoleHandler{t}
}

// Create is mocked Create method of IUserRoleService.Create
func (m *mockUserRoleHandler) Create(input d.UserRoleRequest) (*d.UserRoleResponse, error) {
    // simulate input invalid
    if !input.IsValid() {
        return nil, E.New(E.ErrDataIsInvalid)
    }

    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    // prepare mock data response
    res := ur[0].ConvertToResponse()

    return res, nil
}

// Get is mocked Get method of IUserRoleService.Get
func (m *mockUserRoleHandler) Get(id int) (*d.UserRoleResponse, error) {
    if len(ur) < id {
        return nil, E.New(E.ErrParamIsInvalid)
    }

    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    // prepare mock data response
    res := ur[1].ConvertToResponse()

    return res, nil
}

// Gets is mocked Gets method of IUserRoleService.Gets
func (m *mockUserRoleHandler) Gets() ([]*d.UserRoleResponse, error) {
    // return nil if force error set to true
    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    // convert user.role slice into user.role response dto slice
    var urRes []*d.UserRoleResponse
    for _, ures := range ur {
        urRes = append(urRes, ures.ConvertToResponse())
    }

    return urRes, nil
}

// Update is mocked Update method of IUserRoleService.Update
func (m *mockUserRoleHandler) Update(id int, input d.UserRoleRequest) (*d.UserRoleResponse, error) {
    if len(ur) < id {
        return nil, E.New(E.ErrParamIsInvalid)
    }

    if !input.IsValid() {
        log.Printf("INVALID input: %v\n", input)
        return nil, E.New(E.ErrRequestDataInvalid)
    }

    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    // prepare mock data response
    res := ur[1].ConvertToResponse()

    return res, nil
}

// Delete is mocked Delete method of IUserRoleService.Delete
func (m *mockUserRoleHandler) Delete(id int) (*d.UserRoleResponse, error) {
    if len(ur) < id {
        return nil, E.New(E.ErrParamIsInvalid)
    }

    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    // prepare mock data response
    res := ur[1].ConvertToResponse()

    return res, nil
}

// NewTestUserRoleHandler is function wrapper to get the mock handler of our handler layer
func NewTestUserRoleHandler(t *testing.T) *UserRoleHandler{
    t.Helper()

    // set gin to test mode
    gin.SetMode(gin.TestMode)

    // prepare mock
    mock := NewMockUserRoleHandler(t)
    handler := NewUserRoleHandler(mock)

    // return mocked handler
    return handler
}

// NewTestWriterContext is function wrapper to get http writer and gin test context
func NewTestWriterContext() (*httptest.ResponseRecorder, *gin.Context) {
    writer := httptest.NewRecorder()
    context, _ := gin.CreateTestContext(writer)

    return writer, context
}

// TestUserRoleServiceCreate will test behaviour of UserRoleCreateHandler 
func TestUserRoleCreateHandler(t *testing.T) {
    // prepare the test
    handler := NewTestUserRoleHandler(t)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()

        // prepare mock with ur[0] values
        req := new(d.UserRoleRequest)
        req.RoleName = ur[0].RoleName
        req.Description = ur[0].Description

        uRoleJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uRoleJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.UserRoleCreateHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(ur[0].ConvertToResponse())
        assert.NoError(t, err)

        // validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]))
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success inserting new user.role data") 
    })

    // EXPECT FAIL bind json error. Simulation done by setting wantErr to true 
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
        handler.UserRoleCreateHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusBadRequest, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrDataIsInvalidMsg)
    })

    // EXPECT FAIL insert error. Simulation done by removing role.name field to make 
    // request data invalid
    t.Run("EXPECT FAIL insert record error", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()

        // prepare mock with ur[0] values
        req := new(d.UserRoleRequest)
        req.Description = ur[0].Description

        uRoleJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(uRoleJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method handler call
        handler.UserRoleCreateHandler(context)

        // validation and verification
        assert.Equal(t, http.StatusInternalServerError, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrDataIsInvalidMsg)

    })
}

// TestUserRoleGetHandler will test behaviour of UserRoleGetHandler
func TestUserRoleGetHandler(t *testing.T) {
    // prepare the test
    handler := NewTestUserRoleHandler(t)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
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
        handler.UserRoleGetHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(ur[1].ConvertToResponse())
        assert.NoError(t, err)

        // test validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]))
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success get user.role data")
    })

    // EXPECT FAIL bad param id. Simulated by remove the id param
    t.Run("EXPECT FAIL bad param id", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("GET", "/:id", nil)
        assert.NoError(t, err)

        // actual method call
        handler.UserRoleGetHandler(context)

        assert.Equal(t, http.StatusBadRequest, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrParamIsInvalidMsg)
    })

    // EXPECT FAIL get data error. Simulated by inserting non existing id
    t.Run("EXPECT FAIL get data error", func(t *testing.T) {
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

        // set wantErr to force error return
        wantErr = true

        // actual method call
        handler.UserRoleGetHandler(context)

        // set back wantErr to its default falue
        wantErr = false

        assert.Equal(t, http.StatusInternalServerError, writer.Code)
    })
}

// TestUserRoleGetsHandler will test behaviour of UserRoleGetsHandler
func TestUserRoleGetsHandler(t *testing.T) {
    // prepare the test
    handler := NewTestUserRoleHandler(t)

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
        handler.UserRoleGetsHandler(context)

        // test validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success get user.role data")
    })

    // EXPECT FAIL get data error. Simulated by inserting non existing id
    t.Run("EXPECT FAIL get data error", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("GET", "/", nil)
        assert.NoError(t, err)

        // set wantErr to force error return
        wantErr = true

        // actual method call
        handler.UserRoleGetsHandler(context)

        // set back wantErr to its default falue
        wantErr = false

        assert.Equal(t, http.StatusInternalServerError, writer.Code)
    })
}

// TestUserRoleUpdateHandler will test behaviour of UserRoleUpdateHandler
func TestUserRoleUpdateHandler(t *testing.T) {
    // prepare the test
    handler := NewTestUserRoleHandler(t)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:"1"},
        }

        // prepare mock with ur[0] values
        req := new(d.UserRoleRequest)
        req.RoleName = ur[0].RoleName
        req.Description = ur[0].Description

        uRoleJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("PUT", "/:id", bytes.NewBuffer(uRoleJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method call
        handler.UserRoleUpdateHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(ur[1].ConvertToResponse())
        assert.NoError(t, err)

        // test validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]))
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success update user.role data")
    })

    // EXPECT FAIL bad param id. Simulated by remove the id param
    t.Run("EXPECT FAIL bad param id", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // inject json to request body
        var err error
        context.Request, err = http.NewRequest("PUT", "/:id", nil)
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // actual method call
        handler.UserRoleUpdateHandler(context)

        assert.Equal(t, http.StatusBadRequest, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrParamIsInvalidMsg)
    })

    // EXPECT FAIL inding json error. Simulated by insering invalid body request as input 
    t.Run("EXPECT FAIL binding json error", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:"5"},
        }

        // inject json to request body
        var err error
        context.Request, err = http.NewRequest("PUT", "/:id", nil)
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // set wantErr to force error return
        handler.UserRoleUpdateHandler(context)

        assert.Equal(t, http.StatusBadRequest, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrRequestDataInvalidMsg)
    })

    // EXPECT FAIL update data error. Simulated by forcing error return from the mock
    t.Run("EXPECT FAIL update data error", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:"5"},
        }

        // prepare mock with ur[0] values
        req := new(d.UserRoleRequest)
        req.Description = ur[0].Description

        uRoleJSON, err := json.Marshal(req)
        assert.NoError(t, err)

        // inject json to request body
        context.Request, err = http.NewRequest("PUT", "/:id", bytes.NewBuffer(uRoleJSON))
        assert.NoError(t, err)

        // set content type to json
        context.Request.Header.Add("content-type", "application/json")

        // set wantErr to force error return
        // actual method call
        wantErr = true
        handler.UserRoleUpdateHandler(context)
        wantErr = false

        assert.Equal(t, http.StatusInternalServerError, writer.Code)
    })
}

// TestUserRoleDeleteHandler will test behaviour of UserRoleDeleteHandler
func TestUserRoleDeleteHandler(t *testing.T) {
    // prepare the test
    handler := NewTestUserRoleHandler(t)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:"1"},
        }

        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("DELETE", "/:id", nil)
        assert.NoError(t, err)

        // actual method call
        handler.UserRoleDeletesHandler(context)

        // prepare expected data for test comparison
        want, err := json.Marshal(ur[1].ConvertToResponse())
        assert.NoError(t, err)

        // test validation and verification
        assert.Equal(t, http.StatusOK, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), string(want[:]))
        assert.Contains(t, string(writer.Body.Bytes()[:]), "success delete user.role data")
    })

    // EXPECT FAIL bad param id. Simulated by remove the id param
    t.Run("EXPECT FAIL bad param id", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("DELETE", "/:id", nil)
        assert.NoError(t, err)

        // actual method call
        handler.UserRoleDeletesHandler(context)

        assert.Equal(t, http.StatusBadRequest, writer.Code)
        assert.Contains(t, string(writer.Body.Bytes()[:]), E.ErrParamIsInvalidMsg)
    })

    // EXPECT FAIL get data error. Simulated by inserting non existing id
    t.Run("EXPECT FAIL get data error", func(t *testing.T) {
        // prepare request/ response / gin context
        writer, context := NewTestWriterContext()
        
        // add test context param with key "id"
        context.Params = gin.Params{
            {Key:"id", Value:"1"},
        }

        // inject json to request body
        var err error = nil
        context.Request, err = http.NewRequest("DELETE", "/:id", nil)
        assert.NoError(t, err)

        // set wantErr to force error return
        // actual method call
        wantErr = true
        handler.UserRoleDeletesHandler(context)
        wantErr = false

        assert.Equal(t, http.StatusInternalServerError, writer.Code)
    })
}
