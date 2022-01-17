package domain

import "time"

type UserRoles []UserRole

// IUserRole is 'user.role' interface
type IUserRole interface {
    Close()
    // FindByID(id int) (*UserRole, error)
    // Find() (*UserRoles, error)
    Create(input UserRole) (*UserRole, error)
    // Update(id int, input UserRole) (*UserRole, error)
}

// UserRole is User Role model
type UserRole struct {
    BaseModel
    RoleName string `json:"role_name" binding:"required"`
    Description string `json:"description"`
}

// BeforeCreate() is hook for UserRole before record creation
func (ur UserRole) BeforeCreate() error {
    ur.CreatedAt = time.Now()
    ur.UpdatedAt = time.Now()
    return nil
}

// BeforeUpdate() method is hook for UserRole before update operation
func (ur UserRole) BeforeUpdate() error {
    ur.UpdatedAt = time.Now()
    return nil
}
