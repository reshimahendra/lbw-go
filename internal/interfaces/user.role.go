/*
   packgage interfaces
   user.role.go
   - interface for user.role 
*/
package interfaces

import "github.com/reshimahendra/lbw-go/internal/domain"

// IUserRole is user.role interface for CRUD operation directly
// to the database
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
    Update(id int, input domain.UserRoleRequest) (*domain.UserRole, error)

    // Delete will do 'soft delete' instead of deleting the user record 
    // from the database. Data should be persistant in the database
    Delete(id int) (*domain.UserRole, error)
}

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
