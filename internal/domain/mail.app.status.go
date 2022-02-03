/*
   package domain
   mail.app.status.go
   - containing user mail membership status model, request dto and response dto struct
*/
package domain

// MemberStatus hold membership_status table in database
type MemberStatus struct {
    // ID is user.role id. it is its primary key
    ID          int         `json:"id"`

    // StatusName is MemberStatus status name
    StatusName  string      `json:"status_name"`

    // Description is the short description of the member status
    Description string      `json:"description,omitempty"`
} 

// IsValid is to check whether data MemberStatus record is valid
func (m *MemberStatus) IsValid() bool {
    return m.ID > 0 && m.StatusName != ""
}

