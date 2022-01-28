/*
   package datastore (test)
   - 'user.role' test unit
*/
package datastore 

import (
	"regexp"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var (
    urHeader = []string{"id","role_name","description","created_at","updated_at"}
    ur = []*domain.UserRole{
        {ID : 0, RoleName: "Guest", Description: "Guest role", CreatedAt: time.Now(), UpdatedAt: time.Now()},
        {ID : 1, RoleName: "Superuser", Description: "Superuser role", CreatedAt: time.Now(), UpdatedAt: time.Now()},
        {ID : 2, RoleName: "Admin", Description: "Admin role", CreatedAt: time.Now(), UpdatedAt: time.Now()},
    }
    errUser = domain.UserRole{ID : 3, RoleName: "FAIL", Description: "FAIL role", CreatedAt: time.Now()}
)

// Run will prepare our pgxmock connection interface
func PrepareMock(t *testing.T) pgxmock.PgxPoolIface{
    t.Helper()
    mock, err := pgxmock.NewPool()
    if err != nil {
        t.Errorf("unexpected error occur: %v\n", err)
    }
    defer mock.Close()

    return mock
}

// TestUserRolCreate to test Create metod for user.role
func TestUserRoleCreate(t *testing.T) {
    mock := PrepareMock(t)

    // EXPECT SUCCESS is typical test simulation with expectation that
    // the operation will run normally (successful insert new data and
    // and response with the inserted data)
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleC)).
            WithArgs(ur[0].RoleName,ur[0].Description).
            WillReturnRows(pgxmock.NewRows(urHeader).
                AddRow(ur[0].ID, ur[0].RoleName, ur[0].Description,ur[0].CreatedAt,ur[0].UpdatedAt),
            )

        store := NewUserRoleStore(mock)
        got, err := store.Create(*ur[0])
        
        assert.NoError(t, err)
        assert.Equal(t, ur[0].ID, got.ID)
        assert.Equal(t, ur[0].RoleName, got.RoleName)
    })

    // EXPECT FAIL no row will semulate no row/data empty return error
    t.Run("EXPECT FAIL data is empty", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleC)).
            WithArgs(errUser.RoleName, errUser.Description).
            WillReturnError(pgx.ErrNoRows)

        store := NewUserRoleStore(mock)
        got, err := store.Create(errUser)
        
        assert.Error(t, err)
        assert.Equal(t, E.New(E.ErrDataIsEmpty), err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL database error will semulate database empty return error
    t.Run("EXPECT FAIL database error", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleC)).
            WithArgs(ur[0].RoleName,ur[0].Description).
            WillReturnError(E.New(E.ErrDatabase))

        store := NewUserRoleStore(mock)
        got, err := store.Create(*ur[0])
        
        assert.Error(t, err)
        assert.Equal(t, E.New(E.ErrDatabase), err)
        assert.Nil(t, got)
    })
}

// TestUserRoleGet will test get method of user.role
func TestUserRoleGet(t *testing.T) {
    // prepare mock interface
    mock := PrepareMock(t)

    // EXPECT SUCCESS simulate to get user.role data normally as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleR1)).
            WithArgs(ur[1].ID).
            WillReturnRows(pgxmock.NewRows(urHeader).
                AddRow(ur[1].ID,ur[1].RoleName,ur[1].Description,ur[1].CreatedAt,ur[1].UpdatedAt),
            )

        store := NewUserRoleStore(mock)
        got, err := store.Get(ur[1].ID)
        assert.NoError(t, err)
        assert.Equal(t, ur[1], got)
    })

    // EXPECT FAIL no data simulate to get user.role data with no row return 
    t.Run("EXPECT FAIL data is empty", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleR1)).
            WithArgs(4).
            WillReturnError(pgx.ErrNoRows)

        store := NewUserRoleStore(mock)
        got, err := store.Get(4)
        assert.Error(t, err)
        assert.Equal(t, E.New(E.ErrDataIsEmpty), err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL database error simulate to get user.role data with database error
    // it trigger by the absence of UpdatedAt column
    t.Run("EXPECT FAIL database error", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleR1)).
            WithArgs(errUser.ID).
            WillReturnError(E.New(E.ErrDatabase))

        store := NewUserRoleStore(mock)
        got, err := store.Get(errUser.ID)
        assert.Error(t, err)
        assert.Equal(t, E.New(E.ErrDatabase), err)
        assert.Nil(t, got)
    })
}

// TestUserRoleGets will test behaviour gets method of user.role
func TestUserRoleGets(t *testing.T) {
    // prepare mock interface
    mock := PrepareMock(t)

    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare mock
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleR)).
            WillReturnRows(pgxmock.NewRows(urHeader).
                AddRow(ur[0].ID,ur[1].RoleName,ur[0].Description,ur[0].CreatedAt,ur[0].UpdatedAt).
                AddRow(ur[1].ID,ur[1].RoleName,ur[1].Description,ur[1].CreatedAt,ur[1].UpdatedAt).
                AddRow(ur[2].ID,ur[2].RoleName,ur[2].Description,ur[2].CreatedAt,ur[2].UpdatedAt),
            )

        // actual test method/function call
        store := NewUserRoleStore(mock)
        got, err := store.Gets()

        assert.NoError(t, err)
        assert.Equal(t, ur[1].RoleName, got[1].RoleName)
    })

    // EXPECT FAIL no data simulate to get all user.role data with no row return 
    t.Run("EXPECT FAIL data is empty", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleR)).
            WillReturnError(pgx.ErrNoRows)

        store := NewUserRoleStore(mock)
        got, err := store.Gets()

        assert.Error(t, err)
        assert.Equal(t, E.New(E.ErrDatabase), err)
        assert.Nil(t, got)
    })
}

// TestUserRoleUpdate will test behaviour of user.role Delete method
func TestUserRoleUpdate(t *testing.T) {
    // prepare mock interface
    mock := PrepareMock(t)

    // EXPECT SUCCESS will test normal/ typical behaviour with successfull update result
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare mock
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleU)).
            WithArgs(ur[1].ID,ur[1].RoleName,ur[1].Description).
            WillReturnRows(pgxmock.NewRows(urHeader).
                AddRow(ur[1].ID,ur[1].RoleName,ur[1].Description,ur[1].CreatedAt,ur[1].UpdatedAt),
            )

        // actual method call (method to test)
        store := NewUserRoleStore(mock)
        got, err := store.Update(ur[1].ID, *ur[1])

        // test verification and validation
        assert.NoError(t, err)
        assert.Equal(t, ur[1].ID, got.ID)
    })

    // EXPECT FAIL data is invalid. simulation by removing user.role["role_name"] field
    t.Run("EXPECT FAIL data invalid", func(t *testing.T){
        // prepare mock
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleU)).
            WillReturnError(E.New(E.ErrDataIsInvalid))

        // actual method call (method test)
        store := NewUserRoleStore(mock)
        got, err := store.Update(1, domain.UserRole{Description:"test role fail"})

        // test verification and validation
        assert.Error(t, err)
        assert.Error(t, E.New(E.ErrDataIsInvalid), err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL database error. Will simulate to update user.role data with 
    // database error return. It trigger by the absence of UpdatedAt column
    t.Run("EXPECT FAIL database error", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleU)).
            WillReturnError(E.New(E.ErrDatabase))

        store := NewUserRoleStore(mock)
        got, err := store.Update(4, errUser)

        assert.Error(t, err)
        assert.Equal(t, E.New(E.ErrDatabase), err)
        assert.Nil(t, got)
    })

}


// TestUserRoleDelete will test user.role Delete method behaviour
func TestUserRoleDelete(t *testing.T) {
    // prepare mock interface
    mock := PrepareMock(t)

    // EXPECT SUCCESS will test tipical/normal behaviour of the method
    // which is returning no error / success operation
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // prepare mock
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleD)).
            WithArgs(ur[0].ID).
            WillReturnRows(pgxmock.NewRows(urHeader).
                AddRow(ur[0].ID,ur[0].RoleName,ur[0].Description,ur[0].CreatedAt,ur[0].UpdatedAt),
            )

        // actual method call (method test)
        store := NewUserRoleStore(mock)
        got, err := store.Delete(ur[0].ID)

        want := ur[0]

        // test verification and validation
        assert.NoError(t, err)
        assert.Equal(t, want, got)
    })

    // EXPECT FAIL. simulated by deleting non existing record/ id not found
    t.Run("EXPECT FAIL", func(t *testing.T){
        // prepare mock
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserRoleD)).
            WillReturnError(E.New(E.ErrDatabase))

        // actual method call (method to test)
        store := NewUserRoleStore(mock)
        got, err := store.Delete(8)

        // test verification and validation
        assert.Error(t, err)
        assert.Nil(t, got)
    })

}
