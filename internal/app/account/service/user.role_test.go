/* 
    package service
    user.role_test.go
    - test unit for user.role
*/
package service

import (
	"testing"
	"time"

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

// mockUserRoleService is mocked userRoleService
type mockUserRoleService struct {
    t *testing.T
}

// NewMockUserRoleService is new instance for mockUserRoleService
func NewMockUserRoleService(t *testing.T) *mockUserRoleService{
    return &mockUserRoleService{t: t}
}

// Create is mocked Create method to satisfy IUserRoleStore interface
func (s *mockUserRoleService) Create(input d.UserRole) (*d.UserRole, error) {
    if !input.IsValid() {
        return nil, E.New(E.ErrDataIsInvalid)
    }

    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    // fill empty data with mocked data 
    input.ID = 1
    input.CreatedAt = time.Now()
    input.UpdatedAt = time.Now()

    // return mocked data as return value
    return &input, nil
}

// Get is mocked Get method to satisfy IUserRoleStore interface
func (s *mockUserRoleService) Get(id int) (*d.UserRole, error) {
    if wantErr {
        return nil, E.New(E.ErrDataIsEmpty)
    }

    return ur[id], nil
}

// Gets is mocked Gets method to satisfy IUserRoleStore interface
func (s *mockUserRoleService) Gets() ([]*d.UserRole, error) {
    if wantErr {
        return nil, E.New(E.ErrDataIsEmpty)
    }

    return ur, nil
}

// Update is mocked Update method to satisfy IUserRoleStore interface
func (s *mockUserRoleService) Update(id int, input d.UserRole) (*d.UserRole, error) {
    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    return ur[id], nil
}

// Delete is mocked Delete method to satisfy IUserRoleStore interface
func (s *mockUserRoleService) Delete(id int) (*d.UserRole, error) {
    if wantErr {
        return nil, E.New(E.ErrDatabase)
    }

    return ur[id], nil
}

// TestUserRoleServiceCreate will test "Create" method for user.role service
func TestUserRoleServiceCreate(t *testing.T) {
    // prepare mock and service
    mock := NewMockUserRoleService(t)
    service := NewUserRoleService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        var urReq = new(d.UserRoleRequest)
        urReq.RoleName = ur[0].RoleName
        urReq.Description = ur[0].Description

        got, err := service.Create(*urReq)

        assert.NoError(t, err)
        assert.NotNil(t, got)
    })

    // EXPECT FAIL invalid data will simulated invalid data input with expected 
    // error return is ErrDataIsInvalid. Simulation done by removing 'role_name' field
    t.Run("EXPECT FAIL invalid data", func (t *testing.T) {
        var urReq = new(d.UserRoleRequest)
        urReq.Description = ur[0].Description

        got, err := service.Create(*urReq)

        assert.Error(t, err)
        assert.Nil(t, got)
        // t.Logf("ERROR: %v\n", err)
    })

    // EXPECT FAIL database error will simulated fail inserting record
    // error return is ErrDatabase. Simulation done by triggering 'wantError' variable 
    // so the mocked interface return an error 
    t.Run("EXPECT FAIL database error", func (t *testing.T) {
        // prepare new UserRoleRequest instance
        var urReq = new(d.UserRoleRequest)
        urReq.RoleName = ur[0].RoleName

        // trigger error from the mocked interface
        wantErr = true

        // actual method call (method to test)
        got, err := service.Create(*urReq)

        assert.Error(t, err)
        assert.Nil(t, got)

        // set 'wantErr' value back to default so another test not affected
        wantErr = false
    })
}

// TestUserRoleServiceGet will test behaviour of 'get' method of 'user.role' service
func TestUserRoleServiceGet(t *testing.T) {
    // prepare mock and service
    mock := NewMockUserRoleService(t)
    service := NewUserRoleService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call (method to test)
        got, err := service.Get(ur[0].ID)

        assert.NoError(t, err)
        assert.Equal(t, ur[0].ConvertToResponse(), got)
    })

    // EXPECT FAIL data is empty error will simulated fail getting record
    // error return is ErrDataIsEmpty. Simulation done by triggering 'wantError' variable 
    // so the mocked interface return an error 
    t.Run("EXPECT FAIL database error", func (t *testing.T) {
        // trigger error from the mocked interface
        wantErr = true

        // actual method call (method to test)
        got, err := service.Get(5)

        // test validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)

        // set 'wantErr' value back to default so another test not affected
        wantErr = false
    })
}

// TestUserRoleServiceGets will test behaviour of 'gets' method of 'user.role' service
func TestUserRoleServiceGets(t *testing.T) {
    // prepare mock and service
    mock := NewMockUserRoleService(t)
    service := NewUserRoleService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call (method to test)
        got, err := service.Gets()

        var want []*d.UserRoleResponse
        for _, urRes := range ur {
            want = append(want, urRes.ConvertToResponse())
        }

        assert.NoError(t, err)
        assert.Equal(t, want, got)
    })

    // EXPECT FAIL data is empty error will simulated fail getting record
    // error return is ErrDataIsEmpty. Simulation done by triggering 'wantError' variable 
    // so the mocked interface return an error 
    t.Run("EXPECT FAIL database error", func (t *testing.T) {
        // trigger error from the mocked interface
        wantErr = true

        // actual method call (method to test)
        got, err := service.Gets()

        // test validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)

        // set 'wantErr' value back to default so another test not affected
        wantErr = false
    })
}

// TestUserRoleServiceUpdate will test behaviour of Update method for the user.role service
func TestUserRoleServiceUpdate(t *testing.T) {
    mock := NewMockUserRoleService(t)
    store := NewUserRoleService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare new UserRoleRequest instance
        var urReq = new(d.UserRoleRequest)
        urReq.RoleName = ur[0].RoleName
        urReq.Description = ur[0].Description

        // actual method call (tested method)
        got, err := store.Update(ur[0].ID, *urReq)

        // validation and verification
        assert.NoError(t, err)
        assert.Equal(t, ur[0].ConvertToResponse(), got)
    })

    // EXPECT FAIL data is empty error will simulated fail getting record
    // error return is ErrDataIsEmpty. Simulation done by triggering 'wantError' variable 
    // so the mocked interface return an error 
    t.Run("EXPECT FAIL database error", func (t *testing.T) {
        // trigger error from the mocked interface
        wantErr = true

        // prepare new UserRoleRequest instance
        var urReq = new(d.UserRoleRequest)

        // actual method call (tested method)
        got, err := store.Update(ur[0].ID, *urReq)

        // validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)
        // set 'wantErr' value back to default so another test not affected
        wantErr = false
    })
}

// TestUserRoleServiceDelete will test behaviour of Delete method for the user.role service
func TestUserRoleServiceDelete(t *testing.T) {
    mock := NewMockUserRoleService(t)
    store := NewUserRoleService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call (tested method)
        got, err := store.Delete(ur[0].ID)

        // validation and verification
        assert.NoError(t, err)
        assert.Equal(t, ur[0].ConvertToResponse(), got)
    })

    // EXPECT FAIL data is empty error will simulated fail getting record
    // error return is ErrDataIsEmpty. Simulation done by triggering 'wantError' variable 
    // so the mocked interface return an error 
    t.Run("EXPECT FAIL database error", func (t *testing.T) {
        // trigger error from the mocked interface
        wantErr = true

        // actual method call (tested method)
        got, err := store.Delete(ur[0].ID)

        // validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)
        // set 'wantErr' value back to default so another test not affected
        wantErr = false
    })
}
