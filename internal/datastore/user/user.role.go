/*
    package user
    user.role.go
    - implementing standar CRUD operation for user.role
*/
package user

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/interfaces"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
)

// UserRoleStore is instance wrapper for IDatabase interface
type UserRoleStore struct {
    // db is instance of IDatabase interface
    db interfaces.IDatabase
}

// Close will close the database connection
func (st *UserRoleStore) Close() {
	st.db.Close()
}

// Create will insert new user.role record data into the database
func (st *UserRoleStore) Create(input domain.UserRole) (*domain.UserRole, error) {
    // prepare insert query
    q := `INSERT INTO golang.user_role (role_name,description) VALUES($1,$2) 
          RETURNING id, role_name, description, created_at, updated_at, deleted_at;`

    // execute sql command using QueryRow method
    result := st.db.QueryRow(context.Background(), q, input.RoleName, input.Description) 

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

    // check if error occur while scanning row 
    if err == pgx.ErrNoRows{
        return nil, E.New(E.ErrDataIsEmpty) 
    } else if err != nil {
        return nil, E.New(E.ErrDatabase)
    }

    // return user.role from the scanned variable (ur)
    return ur, nil
}

func (st *UserRoleStore) Get(id int) (*domain.UserRole, error) {
    // prepare sql command to retreive user.role record and execute it 
    q := `SELECT * FROM lbw.user_role WHERE id = $1`
    result := st.db.QueryRow(context.Background(), q, id)

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

    // check for error
    if err == pgx.ErrNoRows{
        return nil, E.New(E.ErrDataIsEmpty) 
    } else if err != nil {
        return nil, E.New(E.ErrDatabase)
    }

    // return user.role from the scanned variable (ur)
    return ur, nil
}

// Gets will get all user.role record from te database
func (st *UserRoleStore) Gets(ctx context.Context) ([]*domain.UserRole, error) {
    q := `SELECT * FROM lbw.user_role`
    result, err := st.db.Query(ctx, q)

    // check for error
    if err == pgx.ErrNoRows{
        return nil, E.New(E.ErrDataIsEmpty) 
    } else if err != nil {
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
            &ur.DeletedAt,
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
