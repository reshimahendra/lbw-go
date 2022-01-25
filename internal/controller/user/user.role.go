package user

import (
	"github.com/reshimahendra/lbw-go/internal/domain"
	// "github.com/reshimahendra/lbw-go/internal/interfaces"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
)
type IUserRole interface {
    // Create will execute sql query insert new user record into database
    Create(input domain.UserRole) (*domain.UserRole, error)

    // Get will execute sql query to get user record from database
    // based on the given id
    Get(id int) (*domain.UserRole, error)

    // Gets will execute sql query to get all user record from database
    Gets() ([]*domain.UserRole, error)

    // Update will execute sql query to update user record
    // based on given input id and input data 
    Update(id int, input domain.UserRole) (*domain.UserRole, error)

    // Delete will do 'soft delete' instead of deleting the user record 
    // from the database. Data should be persistant in the database
    Delete(id int) (*domain.UserRole, error)
}

type UserRoleService struct{
    Store IUserRole
} 

func NewUserRoleService(store IUserRole) *UserRoleService{
    return &UserRoleService{Store: store}
}

func (s *UserRoleService) Create(input domain.UserRoleRequest) (*domain.UserRoleResponse, error) {
    // check if input is invalid
    if !input.IsValid() {
        return nil, E.New(E.ErrDataIsInvalid)
    }

    // send request to create new record to datastore for further process
    result, err := s.Store.Create(*input.ConvertToUserRole())
    if err != nil {
        return nil, err
    }

    // return the user.role create operation result (response) to handler/controller layer
    return result.ConvertToResponse(), nil
}

func (s *UserRoleService) Get(id int) (*domain.UserRoleResponse, error) {
    // send request to datastore to retreive data with id as requested
    result, err := s.Store.Get(id)
    if err != nil {
        return nil, err
    }

    // return the user.role create operation result (response) to handler/controller layer
    return result.ConvertToResponse(), nil
}
