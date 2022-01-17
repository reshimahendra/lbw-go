package interfaces

import (
	"database/sql"
	"log"

	"github.com/reshimahendra/lbw-go/internal/domain"
)

type Datastore struct {
    db *sql.DB
    roleRepo domain.IUserRole
}

func New(db *sql.DB) *Datastore{
    ur, err := NewUserRole(db)
    if err != nil {
        log.Fatalf("unexpected error occur: %v\n", err)
    }
    return &Datastore{db: db, roleRepo: ur}
}
