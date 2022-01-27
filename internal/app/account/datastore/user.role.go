/*
   package datastore 
   user.role.go
   - implementing standar CRUD operation for user.role
*/
package datastore 

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/database"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
)

const (
    // query command for user.role to Create/ insert new record
    sqlUserRoleC = `INSERT INTO public.user_role (role_name,description,updated_at) VALUES($1,$2,CURRENT_TIMESTAMP) 
        RETURNING id, role_name, description, created_at, updated_at;`

    // query command for user.role to get record by its 'id'
    sqlUserRoleR1 = `SELECT * FROM public.user_role WHERE id = $1 AND deleted_at IS NULL`

    // query command for user.role to get all record
    sqlUserRoleR = `SELECT * FROM public.user_role AND deleted_at IS NULL`

    // query command for user.role to update records based on its 'id' and given new record
    sqlUserRoleU = `UPDATE public.user_role SET 
        role_name=$2,description=$3,updated_at=CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL
        RETURNING id, role_name, description, created_at, updated_at;`

    // query to 'soft' delete user.role
    sqlUserRoleD = `UPDATE public.user_role SET deleted_at=CURRENT_TIMESTAMP WHERE id = $1
        RETURNING id, role_name, description, created_at, updated_at;`
)

// IUserRoleStore is user.role interface for CRUD operation directly
// to the database
type IUserRoleStore interface {
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

// UserRoleStore is instance wrapper for IDatabase interface
type UserRoleStore struct {
    // DB is instance of IDatabase interface
    DB database.IDatabase
}

// NewUserRoleStore will create instance of UserRoleStore
func NewUserRoleStore(iDB database.IDatabase) *UserRoleStore {
    return &UserRoleStore{DB: iDB}
}

// Create will insert new user.role record data into the database
func (st *UserRoleStore) Create(input domain.UserRole) (*domain.UserRole, error) {
    // execute sql command using QueryRow method
    result := st.DB.QueryRow(context.Background(), sqlUserRoleC,
        input.RoleName, 
        input.Description,
    ) 

    // prepare new user.role container as a return value and scan the query result
    ur := new(domain.UserRole)
    err := result.Scan(
        &ur.ID,
        &ur.RoleName,
        &ur.Description,
        &ur.CreatedAt,
        &ur.UpdatedAt,
    )

    // check if error occur while scanning row 
    if err == pgx.ErrNoRows{
        return nil, E.New(E.ErrDataIsEmpty) 
    } else if err != nil {
        log.Printf("ERROR: %v", err)
        return nil, E.New(E.ErrDatabase)
    }

    // return user.role from the scanned variable (ur)
    return ur, nil
}

// Get will get user.role record from the database based on its 'id'
func (st *UserRoleStore) Get(id int) (*domain.UserRole, error) {
    // execute sql command to get user.role by its 'id'
    result := st.DB.QueryRow(context.Background(), sqlUserRoleR1, id)

    // prepare new user.role container as a return value and scan the query result
    ur := new(domain.UserRole)
    err := result.Scan(
        &ur.ID,
        &ur.RoleName,
        &ur.Description,
        &ur.CreatedAt,
        &ur.UpdatedAt,
    )

    // check for error
    if err == pgx.ErrNoRows{
        return nil, E.New(E.ErrDataIsEmpty) 
    } else if err != nil {
        return nil, E.New(E.ErrDatabase)
    }

    // return user.role from the scanned variable (ur)
    return ur, nil
}

// Gets will get all user.role record from the database
func (st *UserRoleStore) Gets() ([]*domain.UserRole, error) {
    // execute sql command to get all user.role record
    result, err := st.DB.Query(context.Background(), sqlUserRoleR)

    // check for error
    if err == pgx.ErrNoRows{
        return nil, E.New(E.ErrDatabase)
    }

    defer result.Close()

    // make new variable of user.role slice as a container for scanned result query operation
    urs := make([]*domain.UserRole, 0)
    for result.Next() {
        ur := new(domain.UserRole)
        err := result.Scan(
            &ur.ID,
            &ur.RoleName,
            &ur.Description,
            &ur.CreatedAt,
            &ur.UpdatedAt,
        )

        // make sure no error while scanning rows record data
        if err != nil {
            return nil, E.New(E.ErrDatabase)
        }

        // insert scanned user.role record data (ur) to user.role slice(urs)
        urs = append(urs, ur)
    }

    // return user.role slice(urs) if no error found
    return urs, nil
}

// Update will update user.role record based it 'id' with given new record value
func (st *UserRoleStore) Update(id int, input domain.UserRole) (*domain.UserRole, error) {
    // check whether input is invalid
    if !input.IsValid() {
        return nil, E.New(E.ErrDataIsInvalid) 
    }

    // execute sql command to update user.role by its 'id'  and given 
    // new record values
    result := st.DB.QueryRow(context.Background(), sqlUserRoleU,
        id, input.RoleName, input.Description)
    
    // prepare new user.role container as a return value and scan the query result
    var ur = new(domain.UserRole)
    err := result.Scan(
        &ur.ID,
        &ur.RoleName,
        &ur.Description,
        &ur.CreatedAt,
        &ur.UpdatedAt,
    )

    // check if error occur while scanning record
    if err != nil {
        return nil, E.New(E.ErrDatabase)
    }

    // return scanned user.role data
    return ur, nil
}

// Delete will 'soft' delete user role record data based on its given 'id'
func (st *UserRoleStore) Delete(id int) (*domain.UserRole, error) {
    // execute sql command to 'soft' delete user.role record
    result := st.DB.QueryRow(context.Background(), sqlUserRoleD, id)

    // prepare new user.role container as a return value and scan the query result
    ur := new(domain.UserRole)
    err := result.Scan(
        &ur.ID,
        &ur.RoleName,
        &ur.Description,
        &ur.CreatedAt,
        &ur.UpdatedAt,
        &ur.DeletedAt,
    )

    // check if error occur while scanning record
    if err != nil {
        return nil, err 
    }
    
    return ur, nil
}
