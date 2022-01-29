/*
   package service
   user.status_test.go
   - test unit for user.status
*/
package service

import (
	"testing"

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

// mockUserStatusService is mock to satisfy IUserServiceStore
type mockUserStatusService struct{
    t *testing.T
}

// NewMockUserStatusService will create instance of mockUserStatusService
func NewMockUserStatusService(t *testing.T) *mockUserStatusService{
    return &mockUserStatusService{t}
}

// Get is mock for datastore.(user.status).Get method
func (m *mockUserStatusService) Get(id int) (*d.UserStatus, error) {
    if wantErr {
        return nil, E.New(E.ErrDataIsEmpty)
    }

    return us[id], nil
}

// Gets is mock for datastore.(user.status).Gets method
func (m *mockUserStatusService) Gets() ([]*d.UserStatus, error) {
    if wantErr {
        return nil, E.New(E.ErrDataIsEmpty)
    }

    return us, nil
}

// TestUserStatusServiceGet will test behaviour of Get and Gets method of user.status store
func TestUserStatusServiceGet(t *testing.T) {
    // prepare mock
    mock := NewMockUserStatusService(t)
    service := NewUserStatusService(mock)

    // EXPECT SUCCESS GET will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS GET", func(t *testing.T){
        // actual method test
        got, err := service.Get(us[0].ID)

        assert.NoError(t, err)
        assert.Equal(t, us[0], got)
    })

    // EXPECT FAIL GET data empty record. Simulated by forcing the mock to return error 
    t.Run("EXPECT FAIL GET data empty error", func(t *testing.T){
        // actual method test. call wantErr to force the mock returning error
        wantErr = true
        got, err := service.Get(1)
        wantErr = false

        assert.Error(t, err)
        assert.Nil(t, got)
    })

    // EXPECT SUCCESS GETS will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS GETS", func(t *testing.T){
        // actual method test
        got, err := service.Gets()

        assert.NoError(t, err)
        assert.Equal(t, us, got)
    })

    // EXPECT FAIL GETS data empty. Simulated by forcing the mockk return error empty data
    t.Run("EXPECT SUCCESS GETS", func(t *testing.T){
        // actual method test. call wantErr to force the mock returning error
        wantErr = true
        got, err := service.Gets()
        wantErr = false

        assert.Error(t, err)
        assert.Nil(t, got)
    })
}
