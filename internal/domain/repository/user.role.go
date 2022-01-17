/*
   Package repository
    ur, err := repo.Insert()   Model "User Role"
   TODO:
   * Get record by ID
   * Get All record
   * Insert new record
   * Update record
*/
package repository

import (
	"database/sql"
)

type UserRole struct {
    db *sql.DB
}

func NewUserRole(db *sql.DB) *UserRole {
    return &UserRole{db: db}
}


