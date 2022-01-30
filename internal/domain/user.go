/*
    package domain
    user.go
    - containing user model, request dto and response dto struct
*/
package domain

import (
	"time"

	"github.com/google/uuid"
)

// User is user model
type User struct {
    // ID is the table primary key with uuid type
    ID          uuid.UUID `json:"id"`

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

    // CreatedAt is creation datetime of the record
    CreatedAt time.Time   `json:"created_at"`

    // UpdatedAt is the last updated datetime of the record
    UpdatedAt time.Time   `json:"updated_at"`

    // DeletedAt is the datetime record was deleted ('soft delete')
    DeletedAt time.Time   `json:"deleted_at"`

    // ActivatedAt is account first activation datetime
    ActivatedAt time.Time `json:"activated_at"`
}

// IsValid is to check whether User data is valid
func (u *User) IsValid() bool {
    return  u.ID.String() != "" &&
            u.Username    != "" &&
            u.PassKey     != "" &&
            u.FirstName   != "" &&
            u.Email       != ""
}

// ConvertToResponse will convert User model to response dto format
func (u *User) ConvertToResponse() *UserResponse{
    // return UserResponse data
    return &UserResponse{
        ID        : u.ID,
        Username  : u.Username,
        FirstName : u.FirstName,
        LastName  : u.LastName,
        Email     : u.Email,
        StatusID  : u.StatusID,
        RoleID    : u.RoleID,
        CreatedAt : u.CreatedAt,
        UpdatedAt : u.UpdatedAt,
    }
}

// BeforeInsert will insert required data to User
// func (u *User) BeforeInsert() {
//     u.ID        = uuid.New()
//     u.CreatedAt = time.Now()
//     u.UpdatedAt = time.Now()
// }

// BeforeUpdate will update 'UpdatedAt' field before execute sql command to database
// func (u *User) BeforeUpdate() {
//     u.UpdatedAt = time.Now()
// }

// UserRequest is user request dto
type UserRequest struct {
    // ID is the table primary key with uuid type
    ID          uuid.UUID `json:"id"`

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
}

// IsValid() method will check whether the user request data is validity
func (u *UserRequest) IsValid() bool {
    return  u.ID.String() != "" &&
            u.Username    != "" &&
            u.PassKey     != "" &&
            u.FirstName   != "" &&
            u.Email       != ""
}

// RequestToUser will convert user request dto to User
func (u *UserRequest) RequestToUser() *User{
    return &User{
        ID        : u.ID,
        Username  : u.Username,
        FirstName : u.FirstName,
        LastName  : u.LastName,
        Email     : u.Email,
        PassKey   : u.PassKey,
        StatusID  : u.StatusID,
        RoleID    : u.RoleID,
    }
}

// UserResponse is User response dto
type UserResponse struct {
    // ID is the table primary key with uuid type
    ID        uuid.UUID `json:"id"`

    // Username is the username for the user, value must be unique
    Username  string        `json:"username"`

    // FirstName is the first name of the user
    FirstName string        `json:"first_name"`

    // LastName is the last name for the user
    LastName  string        `json:"last_name"`

    // email is the valid email of the user
    Email     string        `json:"email"`

    // Status is status held by user
    StatusID  int           `json:"status_id"`

    // RoleID is role given to the user on the system
    RoleID    int           `json:"role_id"`

    // CreatedAt is creation datetime of the record
    CreatedAt time.Time     `json:"created_at"`

    // UpdatedAt is the last updated datetime of the record
    UpdatedAt time.Time     `json:"updated_at"`
}

// UserCredential is struct holding user credential data
type UserCredential struct {
    // ID is the table primary key with uuid type
    ID        uuid.UUID `json:"id"`

    // Username is the username for the user, value must be unique
    Username  string        `json:"username"`

    // Status is status held by user
    StatusID  int           `json:"status_id"`

    // PassKey is the password for the account
    PassKey   string    `json:"password"`
}

// IsValid is to ceck whether credential is valid
func (u *UserCredential) IsValid() bool {
    return u.Username != "" || u.PassKey != ""
}
// IsActive is to check whether user credential is active
func (u *UserCredential) IsActive() bool {
    return u.StatusID == 1
}

// NeedActivation is to check whether user credential is not activated yet
func (u *UserCredential) NeedActivation() bool {
    return u.StatusID == 0
}
