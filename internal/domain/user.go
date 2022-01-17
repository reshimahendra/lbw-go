package domain

import (
	"time"

	"github.com/google/uuid"
)

// User is user model
type User struct {
    BaseModelUUID

    Username  string    `json:"username"`
    FirstName string    `json:"first_name"`
    LastName  string    `json:"last_name"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    Active    bool      `json:"active"`
    RoleID    int       `json:"role_id"`
    Role      *UserRole `json:"role"`
}

// BeforeCreate() method is hook on BeforeCreate record on User model
func (u User) BeforeCreate() error {
    u.ID        = uuid.New()
    u.CreatedAt = time.Now()
    u.UpdatedAt = time.Now()

    return nil
}

// BeforeUpdate() method is hook on BeforeUpdate record on User model
func (u User) BeforeUpdate() error {
    u.UpdatedAt = time.Now()

    return nil
}

