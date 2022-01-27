package service

import (
	"testing"

	"github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/database"
)


type MockUserRoleStore struct {
    DB database.IDatabase
}

func NewMockUserRoleStore(iDB database.IDatabase) *MockUserRoleStore{
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
