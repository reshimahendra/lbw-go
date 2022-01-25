package domain

import (
	"time"
)

// UserRole is User Role model
type UserRole struct {
    ID          int         `json:"id"`
    RoleName    string      `json:"role"`
    Description string      `json:"description,omitempty"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
    DeletedAt   time.Time   `json:"deleted_at,omitempty"`
}

// IsValid is to check whether user role data is valid or not
func (ur *UserRole) IsValid() bool{
    return ur.RoleName != ""
}

// ConvertToResponse will convert user.role model to response dto
func (ur *UserRole) ConvertToResponse() *UserRoleResponse{
    return &UserRoleResponse{
        ID          : ur.ID,
        RoleName    : ur.RoleName,
        Description : ur.Description,
        CreatedAt   : ur.CreatedAt,
        UpdatedAt   : ur.UpdatedAt,
    }
}

// UserRoleRequest is user.role request dto
type UserRoleRequest struct {
    RoleName    string  `json:"role"`
    Description string  `json:"description,omitempty"`
}

// ConvertToUserRole will convert user.role request dto to user.role
func (ur *UserRoleRequest) ConvertToUserRole() *UserRole{
    return &UserRole{
        RoleName : ur.RoleName,
        Description : ur.Description,
    }
}

// IsValid is to check whether user.role.request is valid or not
func (ur *UserRoleRequest) IsValid() bool {
    return ur.RoleName != ""
}

// UserRoleResponse is user.role response dto
type UserRoleResponse struct {
    ID          int         `json:"id"`
    RoleName    string      `json:"role"`
    Description string      `json:"description,omitempty"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}
