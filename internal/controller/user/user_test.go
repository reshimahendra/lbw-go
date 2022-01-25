package user

import (
	"testing"

	"github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/interfaces"
)


type MockUserRoleStore struct {
    DB interfaces.IDatabase
}

func NewMockUserRoleStore(iDB interfaces.IDatabase) *MockUserRoleStore{
    return &MockUserRoleStore{DB: iDB}
}

func (m *MockUserRoleStore) Create(input domain.UserRole) (*domain.UserRole, error) {
    return &domain.UserRole{}, nil
}

func (m *MockUserRoleStore) Get(id int) (*domain.UserRole, error) {
    return &domain.UserRole{}, nil
}

func (m *MockUserRoleStore) Gets() ([]*domain.UserRole, error) {
    return []*domain.UserRole{}, nil
}

func (m *MockUserRoleStore) Update(id int, input domain.UserRole) (*domain.UserRole, error) {
    return &domain.UserRole{}, nil
}

func (m *MockUserRoleStore) Delete(id int) (*domain.UserRole, error) {
    return &domain.UserRole{}, nil
}

func UserRoleTestSetup(t *testing.T) {
    t.Helper()

    // store := NewUserRoleStore()
}
