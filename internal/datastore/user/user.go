/*
    package user (datastore)
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
package user

import (
	"context"

	"github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/interfaces"
)

// UserStore interface is instance wrapper for IDatabase interface
type UserStore struct {
    // DB is IDatabase interface instance
    DB interfaces.IDatabase
}

// NewUserStore will create instance of UserStore
func NewUserStore(iDB interfaces.IDatabase) *UserStore {
    return &UserStore{DB: iDB}
} 

// 
func (st *UserStore) Create(input domain.User) (*domain.User, error) {
    // prepare sql command to insert new user
    q := `INSERT INTO users (id,username,firstname,lastname,email,passkey,status_id,role_id,
          created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) ON CONFLICT DO NOTHING
          RETURNING id,username,firstname,lastname,email,passkey,user_status_id,user,role_id,
          created_at,updated_at,activated_at`

    // add required field before insert to the input request
    input.BeforeCreate()

    // execute sql command to insert new user record
    result := st.DB.QueryRow(context.Background(),q, 
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
