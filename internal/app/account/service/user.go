package service 

import (
	"context"

	"github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/app/account/datastore"
)

// IUserService is service layer for user so the handle layer can
// communicate with the datastore layer. User Service layer interface 
// implementing business logic for user operation
type IUserService interface {
    // Create will send user request data to datastore 
    // and request it to process with user insertion operation and 
    // expect to get UserResponse dto from it
    Create(input domain.UserRoleRequest) (*domain.UserRoleResponse, error)

    // Get will make request to datastore to retreive user record based on
    // given id and expect to get UserResponse dto from the operation
    Get(id int) (*domain.UserRoleResponse, error)

    // Gets will make request to datastore to retreive all user data in dto format
    Gets() ([]*domain.UserRoleResponse, error)

    // Update will make request to datastore to update certain record based on its ID
    // with the given new user value
    Update(id int, input domain.UserRoleRequest) (*domain.UserRoleResponse, error)

    // Delete will make request to datastore to do (soft) delete to give user id record
    Delete(id int) (*domain.UserRoleResponse, error)
}

type UserService struct {
    Store datastore.IUserStore
}

func NewUserService(rs datastore.IUserStore) *UserService{
    return &UserService{Store: rs}
}

func (s *UserService) Get(ctx context.Context, id int) (*domain.UserResponse, error) {
    user, err := s.Store.Get(id)
    if err != nil {
        return nil, err
    }
    
    
    return UserToResponse(*user), nil 
}

func UserToResponse(u domain.User) *domain.UserResponse {
    // role, err := New
   return &domain.UserResponse{}
}
