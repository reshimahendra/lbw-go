package interfaces

import (
	"database/sql"
	"fmt"

	"github.com/reshimahendra/lbw-go/internal/domain"
)

type UserRoleRepo struct {
    db *sql.DB
}

func NewUserRole(db *sql.DB) (domain.IUserRole, error){
    return &UserRoleRepo{db: db}, nil
}

func (i *UserRoleRepo) Close() {
	i.db.Close()
}

func (i *UserRoleRepo) Create(input domain.UserRole) (*domain.UserRole, error) {
    var userRole domain.UserRole
    q := `INSERT INTO golang.user_role (role_name,description) VALUES($1,$2) RETURNING id, role_name, description, created_at, updated_at, deleted_at;`
    // q := `INSERT INTO user_role (name, description) VALUES ($1, $2) RETURNING id;`
    // var lastInsertedID int
    uRole := i.db.QueryRow(q, input.RoleName, input.Description) //.Scan(&lastInsertedID)
    if err := uRole.Err(); err != nil {
        return nil, err
    }
    err := uRole.Scan(
        &userRole.ID,
        &userRole.RoleName,
        &userRole.Description,
        &userRole.CreatedAt,
        &userRole.UpdatedAt,
        &userRole.DeletedAt,
    )

    fmt.Print(userRole)

    return &userRole, err
}

