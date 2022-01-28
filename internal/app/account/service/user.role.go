package service

import (
	"github.com/reshimahendra/lbw-go/internal/app/account/datastore"
	"github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
)

// IUserRoleService is service layer for user.role so the handle layer can
// communicate with the datastore/ database layer. user.role Service layer interface
// implementing business logic for user.role operation
type IUserRoleService interface {
    // Create will send user.role request data to datastore 
    // and request it to process with user.role insertion operation and 
    // expect to get UserRoleResponse dto from it
    Create(input domain.UserRoleRequest) (*domain.UserRoleResponse, error)

    // Get will make request to datastore to retreive user.role record based on
    // given id and expect to get UserRoleResponse from the operation
    Get(id int) (*domain.UserRoleResponse, error)

    // Gets will make request to datastore to retreive all user.role data
    Gets() ([]*domain.UserRoleResponse, error)

    // Update will make request to datastore to update certain record based on its ID
    // with the given new user.role value
    Update(id int, input domain.UserRoleRequest) (*domain.UserRoleResponse, error)

    // Delete will make request to datastore to do (soft) delete to give user.role id record
    Delete(id int) (*domain.UserRoleResponse, error)
}


type UserRoleService struct{
    Store datastore.IUserRoleStore
} 

func NewUserRoleService(store datastore.IUserRoleStore) *UserRoleService{
    return &UserRoleService{Store: store}
}

// Create is service layer to send request to datastore to insert new user.role record
// and response with newly inserted user.role data in dto format
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

// Get is service layer to send request to datastore to get user.role record by id
// in user.role response/ dto format
func (s *UserRoleService) Get(id int) (*domain.UserRoleResponse, error) {
    // send request to datastore to retreive data with id as requested
    result, err := s.Store.Get(id)
    if err != nil {
        return nil, err
    }

    // return the user.role create operation result (response) to handler/controller layer
    return result.ConvertToResponse(), nil
}

// Gets is service layer to send request to datastore to retreive all user.role record in
// user.role response (dto) format
func (s *UserRoleService) Gets() ([]*domain.UserRoleResponse, error) {
    // send request to datastore to retreive data with id as requested
    result, err := s.Store.Gets()
    if err != nil {
        return nil, err
    }

    // convert user.role slice into user.role response dto slice
    var urRes []*domain.UserRoleResponse
    for _, ur := range result {
        urRes = append(urRes, ur.ConvertToResponse())
    }

    // return the user.role create operation result (response) to handler/controller layer
    return urRes, nil
}

// Update is service layer to send request to datastore to update certain record based on its id
func (s *UserRoleService) Update(id int, input domain.UserRoleRequest) (*domain.UserRoleResponse, error) {
    // send request to datastore to do update on certain record
    result, err := s.Store.Update(id, *input.ConvertToUserRole())
    if err != nil {
        return nil, err
    }

    return result.ConvertToResponse(), nil
}

// Delete is service layer to send request to datastore to (soft) delete certain record based on its id
func (s *UserRoleService) Delete(id int) (*domain.UserRoleResponse, error) {
    // send request to datastore to do delete on certain record
    result, err := s.Store.Delete(id)
    if err != nil {
        return nil, err
    }

    return result.ConvertToResponse(), nil
}
