/*
    package datastore
    user.go
    - implementing standar CRUD operation for user
    TODO:
        * Get method
        * Gets method
        * Update method
        * Delete method
        * CheckCredential for login/signin operation
        * UserActivation method
        * UserExist method
*/
package datastore

import (
	"context"

	"github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/database"
)

const (
    // prepare sql command to insert new user
    sqlUserC = `INSERT INTO users (id,username,firstname,lastname,email,passkey,status_id,role_id,
          created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) ON CONFLICT DO NOTHING
          RETURNING id,username,firstname,lastname,email,passkey,user_status_id,user,role_id,
          created_at,updated_at,activated_at`

)

// IUserStore is user interface for CRUD operation directly
// to the database
type IUserStore interface {
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

// UserStore interface is instance wrapper for IDatabase interface
type UserStore struct {
    // DB is IDatabase interface instance
    DB database.IDatabase
}

// NewUserStore will create instance of UserStore
func NewUserStore(iDB database.IDatabase) *UserStore {
    return &UserStore{DB: iDB}
} 

// 
func (st *UserStore) Create(input domain.User) (*domain.User, error) {
    // add required field before insert to the input request
    input.BeforeCreate()

    // execute sql command to insert new user record
    result := st.DB.QueryRow(context.Background(),sqlUserC,
        input.ID,
        input.Username,
        input.FirstName,
        input.LastName,
        input.Email,
        input.PassKey,
        input.StatusID,
        input.RoleID,
        input.CreatedAt,
        input.UpdatedAt,
    )

    // prepare variable container to be used as the result query container
    u := new(domain.User)
    err := result.Scan(
        &u.ID,
        &u.Username,
        &u.FirstName,
        &u.LastName,
        &u.Email,
        &u.PassKey,
        &u.StatusID,
        &u.RoleID,
        &u.CreatedAt,
        &u.UpdatedAt,
        &u.ActivatedAt,
    )

    if err != nil {
        return nil, err
    }

    return u, nil
}
