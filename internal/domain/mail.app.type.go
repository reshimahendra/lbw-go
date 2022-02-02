/*
   package domain
   mail.app.type.go
   - containing user mail membership type model, request dto and response dto struct
*/
package domain

import "time"

// MemberMailAppType hold membership_mail_app_type table in database
type MemberMailAppType struct {
    // ID is user.role id. it is its primary key
    ID          int         `json:"id"`

    // TypeName is MemberMailAppType type name
    TypeName    string      `json:"type_name"`

    // Price is MemberMailAppType subscription fee
    Price       float64     `json:"price"`

    // Description is the short description of the member app type 
    Description string      `json:"description,omitempty"`

    // CreatedAt is the record creation datetime
    CreatedAt   time.Time   `json:"created_at"`

    // UpdatedAt is the record last update datetime
    UpdatedAt   time.Time   `json:"updated_at"`

    // DeletedAt is the datetime record has been deleted (soft delete)
    DeletedAt   *time.Time   `json:"deleted_at,omitempty"`
} 

// IsValid is to check whether data MemberMailAppType is valid
func (m *MemberMailAppType) IsValid() bool {
    return m.ID > 0 && m.TypeName != "" && m.Price >= 0
}

// IsDeleted is to check whether data MemberMailAppType already deleted
func (m *MemberMailAppType) IsDeleted() bool {
    return m.DeletedAt == nil
}


