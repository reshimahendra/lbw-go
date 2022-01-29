/*
   package datastore
   user_test.go
   - test unit for user datastore
*/
package datastore

import (
	"regexp"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pashagolub/pgxmock"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var (
    uHeader = []string{"id","username","first_name","last_name","email",
        "status_id","role_id","created_at","updated_at"}

    u = []*d.User{
        {
            ID        : uuid.New(),
            Username  : "leonard",
            FirstName : "Leo",
            LastName  : "Singa",
            Email     : "leo@gmail.com",
            PassKey   : "secret",
            StatusID  : 1,
            RoleID    : 1,
            CreatedAt : time.Now(),
            UpdatedAt : time.Now(),
        },
        {
            ID        : uuid.New(),
            Username  : "jennydoe",
            FirstName : "Jenny",
            LastName  : "Doe",
            Email     : "jenny@gmail.com",
            PassKey   : "secret",
            StatusID  : 1,
            RoleID    : 0,
            CreatedAt : time.Now(),
            UpdatedAt : time.Now(),
        },
    }
)

// TestUserStoreCreate to test Create method for user datastore
func TestUserStoreCreate(t *testing.T) {
    // prepare mock database interface
    mock := PrepareMock(t)
    store := NewUserStore(mock)

    // EXPECT SUCCESS is typical test simulation with expectation that
    // the operation will run normally (successful insert new data and
    // and response with the inserted data)
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserC)).
            WithArgs(u[0].ID,u[0].Username,u[0].FirstName,u[0].LastName,u[0].Email,u[0].PassKey).
            WillReturnRows(pgxmock.NewRows(uHeader).
                AddRow(u[0].ID,u[0].Username,u[0].FirstName,u[0].LastName,u[0].Email,
                u[0].StatusID,u[0].RoleID,u[0].CreatedAt,u[0].UpdatedAt))

        // actual method
        got, err := store.Create(*u[0])

        // validation and verification
        assert.NoError(t, err)
        assert.NotNil(t, got)
        // t.Logf("GOT: %v\nWANT:%v\n",got,u[0])

        // since we not include passkey in the response, we need to add it manually
        // so the 'want' value will equal to 'got' value
        got.PassKey = u[0].PassKey
        assert.Equal(t, u[0], got)
    })

    // EXPECT FAIL data empty error. Simulated by triggering pgx.ErrNoRows
    t.Run("EXPECT FAIL empty data", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserC)).
            WithArgs(u[0].ID,u[0].Username,u[0].FirstName,u[0].LastName,u[0].Email,u[0].PassKey).
            WillReturnError(pgx.ErrNoRows)

        // actual method test
        got, err := store.Create(*u[0])

        // validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)

    })

    // EXPECT FAIL database error. Simulated by triggering error E.ErrDatabase
    t.Run("EXPECT FAIL database error", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserC)).
            WithArgs(u[0].ID,u[0].Username,u[0].FirstName,u[0].LastName,u[0].Email,u[0].PassKey).
            WillReturnError(E.New(E.ErrDatabase))

        // actual method test
        got, err := store.Create(*u[0])

        // validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)
    })
}

// TestUserStoreGet will test Get method of user datastore
func TestUserStoreGet(t *testing.T) {
    // prepare mock and store
    mock := PrepareMock(t)
    store := NewUserStore(mock)

    // EXPECT SUCCESS is typical test simulation with expectation that
    // the operation will run normally (successful insert new data and
    // and response with the inserted data)
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserR1)).
            WithArgs(u[0].ID).
            WillReturnRows(pgxmock.NewRows(uHeader).
                AddRow(u[0].ID,u[0].Username,u[0].FirstName,u[0].LastName,u[0].Email,
                u[0].StatusID,u[0].RoleID,u[0].CreatedAt,u[0].UpdatedAt))

        // actual method
        got, err := store.Get(u[0].ID)

        // validation and verification
        assert.NoError(t, err)
        assert.NotNil(t, got)
    })
    // EXPECT FAIL data empty error. Simulated by triggering pgx.ErrNoRows on mock
    t.Run("EXPECT FAIL data is empty error", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserR1)).
            WithArgs(u[0].ID).
            WillReturnError(pgx.ErrNoRows)

        // actual method
        got, err := store.Get(u[0].ID)

        // validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)
    })


    // EXPECT FAIL database error. Simulated by triggering E.ErrDatabase on mock
    t.Run("EXPECT FAIL database error", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserR1)).
            WithArgs(u[0].ID).
            WillReturnError(E.New(E.ErrDatabase))

        // actual method
        got, err := store.Get(u[0].ID)

        // validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)
    })
}

// TestUserStoreGets will test Gets method of user datastore
func TestUserStoreGets(t *testing.T) {
    // prepare mock and store
    mock := PrepareMock(t)
    store := NewUserStore(mock)

    // EXPECT SUCCESS is typical test simulation with expectation that
    // the operation will run normally (successful insert new data and
    // and response with the inserted data)
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserR)).
            WillReturnRows(pgxmock.NewRows(uHeader).
                AddRow(u[0].ID,u[0].Username,u[0].FirstName,u[0].LastName,u[0].Email,
                u[0].StatusID,u[0].RoleID,u[0].CreatedAt,u[0].UpdatedAt).
                AddRow(u[1].ID,u[1].Username,u[1].FirstName,u[1].LastName,u[1].Email,
                u[1].StatusID,u[1].RoleID,u[1].CreatedAt,u[1].UpdatedAt),
            )

        // actual method test
        got, err := store.Gets()

        // test verification and validation
        assert.NoError(t, err)
        assert.Equal(t, u[0].ID, got[0].ID)
        assert.Equal(t, u[0].Email, got[0].Email)
        assert.Equal(t, u[0].CreatedAt, got[0].CreatedAt)
        assert.Equal(t, u[1].ID, got[1].ID)
        assert.Equal(t, u[1].Email, got[1].Email)
        assert.Equal(t, u[1].CreatedAt, got[1].CreatedAt)
    })

    // EXPECT FAIL data empty error. Simulated by returning error from the mock
    t.Run("EXPECT FAIL data empty error", func(t *testing.T){
        // prepare mock query
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserR)).
            WillReturnError(E.New(E.ErrDataIsEmpty))

        // actual method test
        got, err := store.Gets()

        // test verification and validation
        assert.Error(t, err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL scan data error. Simulated by 
    // mocking ScanAll func of the pgxscan
    t.Run("EXPECT FAIL scan data error", func(t *testing.T){
        // mock inner function pgxscan.ScanAll
        scanAll := scanAllFunc
        scanAllFunc = func(dst interface{}, rows pgx.Rows) error {
            return E.New(E.ErrDatabase)
        }
        defer func() {
            scanAllFunc = scanAll
        }()

        // prepare mock query
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserR)).
            WillReturnRows(pgxmock.NewRows(uHeader).
                AddRow(u[0].ID,u[0].Username,u[0].FirstName,u[0].LastName,u[0].Email,
                u[0].StatusID,u[0].RoleID,u[0].CreatedAt,u[0].UpdatedAt).
                AddRow(u[1].ID,u[1].Username,u[1].FirstName,u[1].LastName,u[1].Email,
                u[1].StatusID,u[1].RoleID,u[1].CreatedAt,u[1].UpdatedAt),
            )

        // actual method test
        got, err := store.Gets()

        // test verification and validation
        assert.Error(t, err)
        assert.Nil(t, got)
    })
}
