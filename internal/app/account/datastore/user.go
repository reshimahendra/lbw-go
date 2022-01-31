/*
   package datastore
   user.go
   - implementing standar CRUD operation for user
   NOTE of method:
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

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/reshimahendra/lbw-go/internal/database"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/logger"
)

const (
    // prepare sql command to insert new user record
    sqlUserC = `INSERT INTO public.users (id,username,firstname,lastname,email,passkey,updated_at,
            user_status_id,user_role_id) VALUES ($1,$2,$3,$4,$5,$6,CURRENT_TIMESTAMP,$7,$8) 
            RETURNING id,username,firstname,lastname,email,status_id,role_id,
            created_at,updated_at`
    sqlUserR1 = `SELECT id,username,firstname,lastname,email,status_id,role_id,
            created_at,updated_at FROM public.users WHERE id = $1 AND deleted_at IS NULL`
    sqlUserR = `SELECT id,username,firstname,lastname,email,status_id,role_id,
            created_at,updated_at FROM public.users WHERE deleted_at IS NULL ORDER BY created_at`
    sqlUserU = `UPDATE public.users SET username=$2,firstname=$3,lastname=$4,email=$5,passkey=$6,
            status_id=$7,role_id=$8,updated_at=CURRENT_TIMESTAMP WHERE id=$1 RETURNING 
            id,username,firstname,lastname,email,status_id,role_id,created_at,updated_at`
    sqlUserD = `UPDATE public.users SET updated_at=CURRENT_TIMESTAMP,deleted_at=CURRENT_TIMESTAMP 
            WHERE id=$1 RETURNING id, username,firstname,lastname,email,status_id,
            role_id,created_at,updated_at`
    sqlCredentialR = `SELECT id,username,passkey,user_status_id FROM public.users WHERE id=$1`
)

var (
    // userScanAll is pgxscan.ScanAll func wrapper for user
    scanAllFunc = pgxscan.ScanAll
)
// IUserStore is user interface for CRUD operation directly
// to the database
type IUserStore interface {
    // Create will execute sql query to insert new user.role record into the database
    Create(input d.User) (*d.User, error)

    // Get will execute sql query to get user.role record from database
    // based on the given id
    Get(id uuid.UUID) (*d.User, error)

    // Gets will execute sql query to get all user.role record from database
    Gets() ([]*d.User, error)

    // Update will execute sql query to update user.role record
    // based on given input id and input data 
    Update(id uuid.UUID, input d.User) (*d.User, error)

    // Delete will do 'soft delete' instead of deleting the user.role record
    // from the database. Data should be persistant in the database
    Delete(id uuid.UUID) (*d.User, error)
}

// UserStore is instance wrapper for IDatabase interface
type UserStore struct {
    // DB is IDatabase interface instance
    DB database.IDatabase
}

// NewUserStore will create instance of UserStore
func NewUserStore(iDB database.IDatabase) *UserStore {
    return &UserStore{DB: iDB}
} 

// Create will create new User record to database
func (st *UserStore) Create(input d.User) (*d.User, error) {
    // execute sql command to insert new user record
    result := st.DB.QueryRow(context.Background(),sqlUserC,
        input.ID,
        input.Username,
        input.Firstname,
        input.Lastname,
        input.Email,
        input.PassKey,
        input.StatusID,
        input.RoleID,
    )

    // prepare variable container to be used as the result query container
    user := new(d.User)
    err := result.Scan(
        &user.ID,
        &user.Username,
        &user.Firstname,
        &user.Lastname,
        &user.Email,
        &user.StatusID,
        &user.RoleID,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    // check if error occur during scan
    if err == pgx.ErrNoRows{
        logger.Errorf("user.create datastore fail: %v", err)
        return nil, E.New(E.ErrDataIsEmpty)
    } else if err != nil {
        logger.Errorf("user.create datastore fail: %v", err)
        return nil, E.New(E.ErrDatabase)
    }

    return user, nil
}

// Get will get user data from database
func (st *UserStore) Get(id uuid.UUID) (*d.User, error) {
    // execute sql command to retreive user record by id
    result := st.DB.QueryRow(context.Background(), sqlUserR1, id)

    // prepare to scan record data
    user := new(d.User)    
    err := result.Scan(
        &user.ID,
        &user.Username,
        &user.Firstname,
        &user.Lastname,
        &user.Email,
        &user.StatusID,
        &user.RoleID,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    // check if error occur during scan
    if err == pgx.ErrNoRows{
        logger.Errorf("user.get datastore fail: %v", err)
        return nil, E.New(E.ErrDataIsEmpty)
    } else if err != nil {
        logger.Errorf("user.get datastore fail: %v", err)
        return nil, E.New(E.ErrDatabase)
    }

    return user, nil
}

// Gets will get all user data from database
func (st *UserStore) Gets() ([]*d.User, error) {
    // execute sql command to retreive user record
    results, err := st.DB.Query(context.Background(), sqlUserR)
    if err != nil {
        logger.Errorf("user.gets datastore fail: %v", err)
        return nil, E.New(E.ErrDataIsEmpty)
    }
    defer results.Close()

    // prepare variable for data scan container
    users := make([]*d.User, 0)
    if err = scanAllFunc(&users, results); err != nil { 
        logger.Errorf("user.gets datastore scan fail: %v", err)
        return nil, E.New(E.ErrDatabase)
    }

    // return user slice
    return users, nil
}

// Update will update user based on given id
func (st *UserStore) Update(id uuid.UUID, input d.User) (*d.User, error) {
    // execute sql command to update user record
    result := st.DB.QueryRow(context.Background(), sqlUserU,
        id,
        input.Username,
        input.Firstname,
        input.Lastname,
        input.Email,
        input.PassKey,
        input.StatusID,
        input.RoleID,
    )

    // prepare to scan record data
    user := new(d.User)    
    err := result.Scan(
        &user.ID,
        &user.Username,
        &user.Firstname,
        &user.Lastname,
        &user.Email,
        &user.StatusID,
        &user.RoleID,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    // check if error occur during scan
    if err == pgx.ErrNoRows{
        logger.Errorf("user.update datastore fail: %v", err)
        return nil, E.New(E.ErrDataIsEmpty)
    } else if err != nil {
        logger.Errorf("user.update datastore fail: %v", err)
        return nil, E.New(E.ErrDatabase)
    }

    return user, nil
}

// Delete will delete user record based on given id
func (st *UserStore) Delete(id uuid.UUID) (*d.User, error) {
    // execute sql command to delete user record
    result := st.DB.QueryRow(context.Background(), sqlUserD, id)

    // prepare to scan record data
    user := new(d.User)    
    err := result.Scan(
        &user.ID,
        &user.Username,
        &user.Firstname,
        &user.Lastname,
        &user.Email,
        &user.StatusID,
        &user.RoleID,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    // check if error occur during scan
    if err == pgx.ErrNoRows{
        logger.Errorf("user.delete datastore fail: %v", err)
        return nil, E.New(E.ErrDataIsEmpty)
    } else if err != nil {
        logger.Errorf("user.delete datastore fail: %v", err)
        return nil, E.New(E.ErrDatabase)
    }

    return user, nil
}

// GetCredential will get user credential data
func GetCredential(DB database.IDatabase, id uuid.UUID) (*d.UserCredential, error) {
    cred := new(d.UserCredential)
    // if err := pgxscan.Get(context.Background(), DB, &cred, sqlCredentialR, id); err != nil{
    //     logger.Errorf("fail get credential data: %v", err)
    //     return nil
    // }
    err := DB.QueryRow(context.Background(), sqlCredentialR, id).Scan(
        &cred.ID,
        &cred.Username,
        &cred.PassKey,
        &cred.StatusID,
    )

    // check if error occur during scan
    if err == pgx.ErrNoRows{
        logger.Errorf("fail get credential data: %v", err)
        return nil, E.New(E.ErrDataIsEmpty)
    } else if err != nil {
        logger.Errorf("fail get credential data: %v", err)
        return nil, E.New(E.ErrDatabase)
    }

    return cred, nil
}
