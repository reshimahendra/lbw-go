/*
   service package
   user.status.go
   - read only service/ business layer for user.status model
*/
package service

import (
	"github.com/reshimahendra/lbw-go/internal/app/account/datastore"
	d "github.com/reshimahendra/lbw-go/internal/domain"
)

// IUserStatusService is service layer for user.status so the handle layer can
// communicate with the datastore/ database layer. user.status Service layer interface
// implementing business logic for user.status operation (read only)
type IUserStatusService interface{
    // Get will make request to datastore to retreive user.status record based on
    // given id 
    Get(id int) (*d.UserStatus, error)

    // Gets will make request to datastore to retreive all user.status data
    Gets() ([]*d.UserStatus, error)
}

// UserStatusService is type wrapper for interface IUserStatusStore
type UserStatusService struct{
    Store datastore.IUserStatusStore
} 

// NewUserStatusService will create new instance for UserStatusService
func NewUserStatusService(store datastore.IUserStatusStore) *UserStatusService{
    return &UserStatusService{Store: store}
}

// Get is service layer to send request to datastore to get user.status record by id
func (s *UserStatusService) Get(id int) (*d.UserStatus,  error) {
    return s.Store.Get(id)
}

// Gets is service layer to send request to datastore to get all user.status record
func (s *UserStatusService) Gets() ([]*d.UserStatus,  error) {
    return s.Store.Gets()
}
