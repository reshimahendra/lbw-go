/*
   service package
   user.go
   - service/ business layer for user model
*/
package service

import (
	"github.com/google/uuid"
	ds "github.com/reshimahendra/lbw-go/internal/app/account/datastore"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/helper"
	"github.com/reshimahendra/lbw-go/internal/pkg/logger"
)

var (
    // generateHashPassFunc is func instance of helper.HashPassword
    // it will be used to mock the inner func on test
    generateHashPassFunc = helper.HashPassword

    // checkPassHashFunc is func instance of helper.CheckPasswordHash
    checkPassHashFunc = helper.CheckPasswordHash
)

// IUserService is service layer for user so the handle layer can
// communicate with the datastore layer. User Service layer interface
// implementing business logic for user operation
type IUserService interface {
    // Create will send user request data to datastore 
    // and request it to process with user insertion operation and 
    // expect to get UserResponse dto from it
    Create(input d.UserRequest) (*d.UserResponse, error)

    // Get will make request to datastore to retreive user record based on
    // given id and expect to get UserResponse dto from the operation
    Get(id string) (*d.UserResponse, error)

    // GetByEmail will make request to datastore to get user credential data by its username
    GetByEmail(email string) (*d.UserCredential, error) 

    // Gets will make request to datastore to retreive all user data in dto format
    Gets() ([]*d.UserResponse, error)

    // Update will make request to datastore to update certain record based on its ID
    // with the given new user value
    Update(id string, input d.UserRequest) (*d.UserResponse, error)

    // Delete will make request to datastore to do (soft) delete to give user id record
    Delete(id string) (*d.UserResponse, error)

    // GetCredential will make request to datastore to get user credential data
    GetCredential(username,passkey string) (*d.UserCredential, error) 

    // IsUserExist will make request to datastore to check whether username/ email
    // is already exist
    IsUserExist(username,email string) bool
}

// UserService is instance wrapper for IUserStore interface
type UserService struct {
    Store ds.IUserStore
}

// NewUserService is new instance of UserService
func NewUserService(rs ds.IUserStore) *UserService{
    return &UserService{Store: rs}
}

// Create will send request to datastore to insert new user record
func (s *UserService) Create(input d.UserRequest) (*d.UserResponse, error) {
    // create new uuid
    input.ID = uuid.New()

    // check if input data is invalid
    if !input.IsValid() {
        err := E.New(E.ErrDataIsInvalid)
        logger.Errorf("%v", err)
        return nil, err
    }

    // check if email is valid
    if !helper.EmailIsValid(input.Email) {
        err := E.New(E.ErrEmailIsInvalid)
        logger.Errorf("%v", err)
        return nil, err
    }

    // generate hashed passkey
    passKey, err := generateHashPassFunc(input.PassKey)
    if err != nil {
        logger.Errorf("generate passkey fail: %v", err)
        return nil, err
    }
    input.PassKey = passKey

    // send request to datastore to insert new user record
    user, err := s.Store.Create(*input.RequestToUser())
    if err != nil {
        return nil, err
    }

    // if no error found, send back response data to handler layer
    return user.ConvertToResponse(), nil
}

// Get will send request to user datastore to retreive user record with given id
func (s *UserService) Get(id string) (*d.UserResponse, error) {
    // send request to datastore to get record
    user, err := s.Store.Get(*ParseUUID(id))
    if err != nil {
        return nil, err
    }

    // return user.response to handler layer
    return user.ConvertToResponse(), nil 
}

// Gets will send request to user datastore to retreive all user record
func (s *UserService) Gets() ([]*d.UserResponse, error) {
    // send request to datastore to get record
    users, err := s.Store.Gets()
    if err != nil {
        return nil, err
    }

    // make new instance of user response as container of the 
    // converting result from user to user.response dto
    uRes := make([]*d.UserResponse, 0)
    for _, u := range users {
        uRes = append(uRes, u.ConvertToResponse())
    }

    // return user.response slice to handler layer
    return uRes, nil 
}

// Update will send request to user datastore to update user record by given user id
func (s *UserService) Update(id string, input d.UserRequest) (*d.UserResponse, error) {
    // check if input data is invalid
    if !input.IsValid() {
        err := E.New(E.ErrDataIsInvalid)
        logger.Errorf("%v", err)
        return nil, err
    }

    // check if email if its valid
    if !helper.EmailIsValid(input.Email) {
        err := E.New(E.ErrEmailIsInvalid)
        logger.Errorf("%v", err)
        return nil, err
    }

    // parse id string to UUID
    userUUID := ParseUUID(id)

    // get user record as comparison data
    user, err := s.Store.Get(*userUUID)
    if err != nil {
        return nil, E.New(E.ErrGettingData)
    }
    
    // check if password change
    if !checkPassHashFunc(input.PassKey, user.PassKey) {
        // it mean password has changed
        newPass, err := generateHashPassFunc(input.PassKey)
        if err != nil {
            logger.Errorf("generate passkey fail: %v", err)
            return nil, err
        }

        // update input passskey with hashed key
        input.PassKey = newPass
    } else {
        // do not update passkey and use the old hashed key
        input.PassKey = user.PassKey
    }

    // update user data
    updatedUser, err := s.Store.Update(*userUUID, *input.RequestToUser())
    if err != nil {
        return nil, err
    }

    // return response to handler layer
    return updatedUser.ConvertToResponse(), nil
}

// Delete will  send request to user datastore to 'soft' delete user record by given user id
func (s *UserService) Delete(id string) (*d.UserResponse, error) {
    // delete user data
    user, err := s.Store.Delete(*ParseUUID(id))
    if err != nil {
        return nil, err
    }

    // return response to handler layer
    return user.ConvertToResponse(), nil
}

// GetByEmail will send request to user datastore to get user credential data
// based on its email
func (s *UserService) GetByEmail(email string) (*d.UserCredential, error) {
    // get user credential data
    cred, err := s.Store.GetByEmail(email)
    if err != nil {
        return nil, err
    }

    return cred, nil
}

// IsUserExist will send request to datastore to check whether username or email

// GetCredential will send request to user datastore to get user credential data
func (s *UserService) GetCredential(username,passkey string) (*d.UserCredential, error) {
    // get user credential data
    cred, err := s.Store.GetCredential(username, passkey)
    if err != nil {
        return nil, err
    }

    return cred, nil
}

// IsUserExist will send request to datastore to check whether username or email
// is already exist
func (s *UserService) IsUserExist(username, email string) bool {
    if !helper.EmailIsValid(email) {
        logger.Errorf("IsUserExist input email invalid: %s", email)
        return false
    }

    found, _ := s.Store.IsUserExist(username, email)

    return found
}

// ParseUUID is wrapper and simplified version of func uuid.Parse(...) 
func ParseUUID(s string) *uuid.UUID{
    // parse uuid string into uuid
    aUUID, err := uuid.Parse(s)
    if err != nil {
        logger.Errorf("fail parse uuid from string: %v", err)
        return nil
    }

    return &aUUID
}
