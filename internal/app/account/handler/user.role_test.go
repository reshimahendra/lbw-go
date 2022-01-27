package handler

import (
	"bytes"
	"encoding/json"
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

type mockUserRoleHandler struct {
    t *testing.T
}

func NewMockUserRoleHandler(t *testing.T) *mockUserRoleHandler{
    return &mockUserRoleHandler{t}
}

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

func (m *mockUserRoleHandler) Get(id int) (*d.UserRoleResponse, error) {
    return nil, nil
}

func (m *mockUserRoleHandler) Gets() ([]*d.UserRoleResponse, error) {
    return nil, nil
}

func (m *mockUserRoleHandler) Update(id int, input d.UserRoleRequest) (*d.UserRoleResponse, error) {
    return nil, nil
}

func (m *mockUserRoleHandler) Delete(id int) (*d.UserRoleResponse, error) {
    return nil, nil
}

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

func NewTestWriterContext() (*httptest.ResponseRecorder, *gin.Context) {
    writer := httptest.NewRecorder()
    context, _ := gin.CreateTestContext(writer)

    return writer, context
}

// TestUserRoleServiceCreate will test "Create" method for user.role service
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

        // prepare mock with ur[0] values
        // req := new(d.UserRoleRequest)
        // req.RoleName = ur[0].RoleName
        // req.Description = ur[0].Description
        //
        // uRoleJSON, err := json.Marshal(req)
        // assert.NoError(t, err)

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
