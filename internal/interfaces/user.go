/*
   packgage interfaces
   user.go
   - interfaces for user model
*/
package interfaces

import "github.com/reshimahendra/lbw-go/internal/domain"

// IUser is user interface for CRUD operation
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
