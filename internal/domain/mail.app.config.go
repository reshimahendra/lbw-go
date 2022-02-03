/*
   package domain
   mail.app.config.go
   - containing user membership mail app configuration model, request dto and response dto struct
*/
package domain

import (
	"time"

	"github.com/google/uuid"
)

// MemberMailAppConfig hold membership_mail_app_config table in database
type MemberMailAppConfig struct {
    // ID is primary key and foreign key of MemberMailAppConfig which refer to table User
    ID                 uuid.UUID    `json:"id"`

    // CfgID is config id, primary key of MemberMailAppConfig with auto increment value
    // based on record count
    CfgID              int          `json:"cfg_id"`

    // ConfigName is configuration name of the record
    ConfigName         string       `json:"config_name"`

    // DefaultConfig is flag to know whether the record is default config or not
    DefaultConfig      bool         `json:"default_config"`

    // SmtpServer is smtp server of the configuration
    SmtpServer         string       `json:"smtp_server"`

    // SmtpPort is smtp port of the configuration
    SmtpPort           string       `json:"smtp_port"`

    // SmtpUsername is smtp username of the configuration
    SmtpUsername       string       `json:"smtp_username"`

    // SmtpPassword is smtp password of the configuration
    SmtpPassword       string       `json:"smtp_password"`

    // SmtpSenderEmail is sender mail of the configuration
    SmtpSenderEmail    string       `json:"smtp_sender_email"`

    // SmtpSenderIdentity is sender identity that shown on 'sent by' header
    SmtpSenderIdentity string       `json:"smtp_sender_identity"`

    // ActiveStatus is flag to know whether this configuration active or not
    ActiveStatus       bool         `json:"active_status"`

    // CreatedAt is datetime the configuration record created
    CreatedAt          time.Time    `json:"created_at"`

    // UpdatedAt is date time the configuration record last updated
    UpdatedAt          time.Time    `json:"updated_at"`

    // DeletedAt is datetime the configuration is deleted (soft delete)
    DeletedAt          *time.Time   `json:"deleted_at"`
}

// IsValid will check whether MemberMailAppConfig record is valid
func (m *MemberMailAppConfig) IsValid() bool {
    return m.CfgID > 0 && 
        m.ConfigName != "" &&
        m.SmtpServer != "" &&
        m.SmtpPort != "" &&
        m.SmtpUsername != "" &&
        m.SmtpPassword != ""
}


