/*
   package interfaces
   - 'user.role' test unit
*/
package interfaces

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
    sqldb, mock, err := sqlmock.New()
    if err != nil {
        log.Fatalf("unexpected error occur: %v\n", err)
    }

    return sqldb, mock
}

func TestInsert(t *testing.T) {
    db, mock := NewMock()

    // t.Logf("SQLDB MOCK: %v", sqldb)
    urTest := domain.UserRole{
        RoleName: "role1",
        Description: "role description test",
    }

    q := `INSERT INTO golang.user_role (role_name,description) VALUES ($1,$2) RETURNING id;`
    mock.ExpectBegin()
    mock.ExpectExec(regexp.QuoteMeta(q)).
        WithArgs(urTest.RoleName, urTest.Description).
        WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()
    
    datastore := New(db)
    got, err := datastore.roleRepo.Create(urTest)
    // db := domain.DB()
    
    // repo := &UserRole{db: sqldb}
    // assert.NotNil(t, repo)
    // got, err := repo.Insert(urTest)
    assert.Error(t, err)
    assert.Nil(t, got)
    t.Logf("Mock: %v, User Role: %v\n", mock, got)
}
