package domain

import (
	"time"

	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
)

// UserRole is User Role model
type UserRole struct {
    BaseModel
    RoleName    string  `json:"role"`
    Description string  `json:"description,omitempty"`
}

// BeforeCreate() is hook for UserRole before record creation
func (ur *UserRole) BeforeCreate() error {
    ur.CreatedAt = time.Now()
    ur.UpdatedAt = time.Now()
    return nil
}

// BeforeUpdate() method is hook for UserRole before update operation
func (ur *UserRole) BeforeUpdate() error {
    ur.UpdatedAt = time.Now()
    return nil
}

// IsValid is to check whether user role data is valid or not
func (ur *UserRole) IsValid() error {
    if ur.RoleName == "" {
        return E.New(E.ErrDataIsEmpty) 
    }

    return nil
}

// ConvertToResponse will convert user.role model to response dto
func (ur *UserRole) ConvertToResponse() *UserRoleResponse{
    return &UserRoleResponse{
        ID          : ur.ID,
        RoleName    : ur.RoleName,
        Description : ur.Description,
    }
}

// UserRoleRequest is user.role request dto
type UserRoleRequest struct {
    RoleName    string  `json:"role"`
    Description string  `json:"description"`
}

// IsValid is to check whether user.role.request is valid or not
func (ur *UserRoleRequest) IsValid() bool {
    return ur.RoleName != ""
}

// UserRoleResponse is user.role response dto
type UserRoleResponse struct {
    ID          int     `json:"id"`
    RoleName    string  `json:"role"`
    Description string  `json:"description,omitempty"`
}
