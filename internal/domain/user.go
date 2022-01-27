/*
    package domain
    user.go
    - containing user model, request dto and response dto struct
*/
package domain

import (
	"time"

	"github.com/google/uuid"
    // E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
)

// UserStatus is model for status of the user
type UserStatus struct {
    // ID is user status id which is its primary key
    ID          int     `json:"id"`

    // StatusName is the name of the status
    StatusName  string  `json:"status"`

    // Description is the short description of the status
    Description string  `json:"description"`
}

// User is user model
type User struct {
    /// BaseModelUUID will embed BaseModelUUID model to user model
    BaseModelUUID

    // Username is the username for the user, value must be unique
    Username    string    `json:"username"`

    // FirstName is the first name of the user
    FirstName   string    `json:"first_name"`

    // LastName is the last name for the user
    LastName    string    `json:"last_name"`

    // email is the valid email of the user
    Email       string    `json:"email"`

    // PassKey is the password for the account
    PassKey     string    `json:"password"`

    // StatusID is id of status held by user
    // ("0=inactive", "1=active", "2=suspended", "3=banned")
    StatusID    int       `json:"status_id"`

    // RoleID is role given to the user on the system
    RoleID      int       `json:"role_id"`

    ActivatedAt time.Time `json:"activated_at"`
}

// BeforeCreate() method is hook on BeforeCreate record on User model
func (u *User) BeforeCreate() {
    // assign user.id with new uuid value
    u.ID        = uuid.New()

    // assign user.CreatedAt value to current date time
    u.CreatedAt = time.Now()

    // assign user.UpdatedAt value to current date time
    u.UpdatedAt = time.Now()
}

// BeforeUpdate() method is hook on BeforeUpdate record on User model
func (u *User) BeforeUpdate() {
    // assign user.UpdatedAt value to current date time
    u.UpdatedAt = time.Now()
}

// IsValid() method will check whether the user data is valid or not
func (u *User) IsValid() bool {
    return u.ID.String() != "" &&
        u.Username != "" &&
        u.PassKey != "" &&
        u.FirstName != "" &&
        u.Email != ""
}

// func (u *User) ConvertToResponse() (*UserResponse, error) {
//     if !u.IsValid {
//         err := E.New(E.ErrDataIsInvalid)
//         return nil, err
//     }
//
//     return &UserResponse{
//         ID : u.ID,
//         Username : u.Username,
//         FirstName : u.FirstName,
//         LastName : u.LastName,
//         Email : u.Email,
//         StatusID : u.StatusID,
//         RoleID : u.RoleID,
//
//     }
// }

type UserResponse struct {
    /// BaseModelUUID will embed BaseModelUUID model to user model
    BaseModelUUID

    // Username is the username for the user, value must be unique
    Username  string        `json:"username"`

    // FirstName is the first name of the user
    FirstName string        `json:"first_name"`

    // LastName is the last name for the user
    LastName  string        `json:"last_name"`

    // email is the valid email of the user
    Email     string        `json:"email"`

    // Status is status held by user
    Status    *UserStatus   `json:"status,omitempty"`

    // RoleID is role given to the user on the system
    Role      *UserRole     `json:"role,omitempty"`
}
