/*
    Package domain
    - Base model
*/
package domain

import (
    "time"
    "github.com/google/uuid"
)

// BaseModel is base model with standard autoincrement primary key
type BaseModel struct {

    ID        int       `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt time.Time `json:"deleted_at"`
}

// BaseModelUUID is base model with uuid primary key
type BaseModelUUID struct{

    ID        uuid.UUID `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt time.Time `json:"deleted_at"`
}
