/*
   package domain
   mail.app.go
   - containing user membership mail app model, request dto and response dto struct
*/
package domain

import (
	"time"

	"github.com/google/uuid"
)

// MemberMailApp hold membership_mail_app table in database
type MemberMailApp struct {
    // ID is MemberMailApp primary key as well as 1:1 foreign key to User
    ID             uuid.UUID    `json:"id"`

    // TypeID is associated id of membership.mail.ap.type
    TypeID         int          `json:"type_id"`

    // StatusID is associated id of membership.status
    StatusID       int          `json:"status_id"`

    // Price is the price of the mail.app service
    Price          float64      `json:"price"`

    // LastPaidAmount is the last paid amount for using the service
    LastPaidAmount float64      `json:"last_paid_amount"`

    // LastPaidAt is the last paid date and time
    LastPaidAt     time.Time    `json:"last_paid_at"`

    // SubscribedAt is the datetime User first register/ subscribe to the service
    SubscribedAt   time.Time    `json:"subscribed_at"`

    // UpdatedAt is datetime of the record last updated
    UpdatedAt      time.Time    `json:"updated_at"`
}

// IsValid will check whether MemberMailApp record is valid
func (m *MemberMailApp) IsValid() bool {
    return m.TypeID >= 0 && m.StatusID >= 0 && m.Price >= 0 && m.LastPaidAmount >= 0
}
