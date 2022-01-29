/*
   datastore package
   user.status.go
   - read only persistent/ datastore layer for user.status model
*/
package datastore

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/reshimahendra/lbw-go/internal/database"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
)

const (
    sqlUserStatusR1 = `SELECT id,status_name,description FROM public.user_status WHERE id = $1`
    sqlUserStatusR  = `SELECT id,status_name,description FROM public.user_status ORDER BY id`
)


// IUserStatusStores is user.status interface for 'Read' / 'Get' operation directly
// to the database
type IUserStatusStore interface{
    // Get will execute sql query to get user record from database
    // based on the given id
    Get(id int) (*d.UserStatus, error)

    // Gets will execute sql query to get all user record from database
    Gets() ([]*d.UserStatus, error)
}

// UserStatusStore is instance wrapper for IDatabase interface
type UserStatusStore struct {
    // DB is instance of IDatabase interface
    DB database.IDatabase
}

// NewUseStatusStore will create instance of UserStatusStore
func NewUserStatusStore(iDB database.IDatabase) *UserStatusStore {
    // return UserStatusStore
    return &UserStatusStore{DB: iDB}
}

// Get will get user.status from the database based on its 'id'
func (s *UserStatusStore) Get(id int) (*d.UserStatus, error) {
    // execute sql command to get user.status by its 'id'
    result := s.DB.QueryRow(context.Background(), sqlUserStatusR1, id)

    // prepare to scan the result query/ rows and place it into our new variable (us)
    us := new(d.UserStatus)
    err := result.Scan(
        &us.ID,
        &us.StatusName,
        &us.Description,
    )

    // check if error occur while scanning row 
    if err == pgx.ErrNoRows{
        return nil, E.New(E.ErrDataIsEmpty) 
    } else if err != nil {
        return nil, E.New(E.ErrDatabase)
    }

    // return user.status from the scanned variable (ur)
    return us, nil
}

// Gets will get all user.status record from the database
func (s *UserStatusStore) Gets() ([]*d.UserStatus, error) {
    // execute sql command to get all user.status
    results, err := s.DB.Query(context.Background(), sqlUserStatusR)
    if err != nil {
        return nil, err
    }
    defer results.Close()

    // prepare to scan the result query/ rows and place it into our new variable (us)
    // scanAllFunc is locate at gatastore - user.go
    var uSts []*d.UserStatus
    if err = scanAllFunc(&uSts, results); err != nil {
        return nil, err
    }

    return uSts, nil
}
