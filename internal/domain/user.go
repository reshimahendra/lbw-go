package domain

import (
	"time"

	"github.com/google/uuid"
)

// UserStatus is model for status of the user
type UserStatus struct {
    ID          int     `json:"id"`
    StatusName  string  `json:"status"`
    Description string  `json:"description"`
}

// User is user model
type User struct {
    /// BaseModelUUID will embed BaseModelUUID model to user model
    BaseModelUUID

    // Username is the username for the user, value must be unique
    Username  string    `json:"username"`

    // FirstName is the first name of the user
    FirstName string    `json:"first_name"`

    // LastName is the last name for the user
    LastName  string    `json:"last_name"`

    // email is the valid email of the user
    Email     string    `json:"email"`

    // PassKey is the password for the account
    PassKey   string    `json:"password"`

    // StatusID is id of status held by user
    // ("0=inactive", "1=active", "2=suspended", "3=banned")
    StatusID  int       `json:"status_id"`

    // RoleID is role given to the user on the system
    RoleID    int       `json:"role_id"`
}

// BeforeCreate() method is hook on BeforeCreate record on User model
func (u *User) BeforeCreate() error {
    // assign user.id with new uuid value
    u.ID        = uuid.New()

    // assign user.CreatedAt value to current date time
    u.CreatedAt = time.Now()

    // assign user.UpdatedAt value to current date time
    u.UpdatedAt = time.Now()

    return nil
}

// BeforeUpdate() method is hook on BeforeUpdate record on User model
func (u *User) BeforeUpdate() error {
    // assign user.UpdatedAt value to current date time
    u.UpdatedAt = time.Now()

    return nil
}

// IsValid() method will check whether the user data is valid or not
func (u *User) IsValid() bool {
    return u.ID.String() != "" &&
        u.Username != "" &&
        u.PassKey != "" &&
        u.FirstName != "" &&
        u.Email != ""
}
