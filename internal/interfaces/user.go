/*
   packgage interfaces
   user.go
   - interfaces for user model
*/
package interfaces

import "github.com/reshimahendra/lbw-go/internal/domain"

// IUser is user interface for CRUD operation directly
// to the database
type IUser interface {
    // Create will execute sql query to insert new user.role record into the database
    Create(input domain.User) (*domain.UserRole, error)

    // Get will execute sql query to get user.role record from database
    // based on the given id
    Get(id int) (*domain.User, error)

    // Gets will execute sql query to get all user.role record from database
    Gets() ([]*domain.User, error)

    // Update will execute sql query to update user.role record
    // based on given input id and input data 
    Update(id string, input domain.User) (*domain.UserRole, error)

    // Delete will do 'soft delete' instead of deleting the user.role record
    // from the database. Data should be persistant in the database
    Delete(id int) (*domain.User, error)
}

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
