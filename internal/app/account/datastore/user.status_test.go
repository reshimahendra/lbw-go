/*
   package datastore (test)
   user.status.go
   - user.status test unit
*/
package datastore

import (
	"regexp"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/pashagolub/pgxmock"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var (
    userStatusHeader = []string{"id","status_name","description"}
    us = []*d.UserStatus{
        {ID:0, StatusName:"inactive",Description:"inactive member"},
        {ID:1, StatusName:"active",Description:"active member"},
        {ID:2, StatusName:"vip",Description:"vip member"},
    }
)

// TestUserStatusGet will test get method of user.status store
func TestUserStatusGet(t *testing.T) {
    // prepare database mock, this func is available on 'user.status_test.go'
    mock := PrepareMock(t)
    store := NewUserStatusStore(mock)

    // EXPECT SUCCESS is typical test simulation with expectation that
    // the operation will run normally (successful get user.status record)
    t.Run("EXPECT SUCCESS GET", func(t *testing.T){
        // prepare mock query
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserStatusR1)).
            WithArgs(us[1].ID).
            WillReturnRows(pgxmock.NewRows(userStatusHeader).
                AddRow(us[1].ID, us[1].StatusName,us[1].Description),
            )

        // actual method test
        got, err := store.Get(us[1].ID)
        assert.NoError(t, err)
        assert.Equal(t, us[1], got)
    })

    // EXPECT FAIL GET no data simulate to get user.status data with no row return 
    t.Run("EXPECT FAIL GET data is empty", func(t *testing.T){
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserStatusR1)).
            WithArgs(4).
            WillReturnError(pgx.ErrNoRows)

        got, err := store.Get(4)
        assert.Error(t, err)
        assert.Equal(t, E.New(E.ErrDataIsEmpty), err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL GET retreive data error. Simulated by returning error from the mock
    t.Run("EXPECT FAIL GET retreive data error", func(t *testing.T){
        // prepare mock query
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserStatusR1)).
            WithArgs(us[1].ID).
            WillReturnError(E.New(E.ErrDataIsInvalid))

        // actual method test
        got, err := store.Get(us[1].ID)
        assert.Error(t, err)
        assert.Nil(t, got)
    })

    // EXPECT SUCCESS is typical test simulation with expectation that
    // the operation will run normally (successful get user.status record)
    t.Run("EXPECT SUCCESS GETS", func(t *testing.T){
        // prepare mock query
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserStatusR)).
            WillReturnRows(pgxmock.NewRows(userStatusHeader).
                AddRow(us[0].ID, us[0].StatusName,us[0].Description).
                AddRow(us[1].ID, us[1].StatusName,us[1].Description).
                AddRow(us[2].ID, us[2].StatusName,us[2].Description),
            )

        // actual method test
        got, err := store.Gets()
        assert.NoError(t, err)
        assert.Equal(t, us, got)
    })

    // EXPECT FAIL GETS data empty error. Simulated by returning error from the mock
    t.Run("EXPECT FAIL GETS data empty error", func(t *testing.T){
        // prepare mock query
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserStatusR)).
            WillReturnError(E.New(E.ErrDataIsEmpty))

        // actual method test
        got, err := store.Gets()
        assert.Error(t, err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL GETS scan data error. Simulated by 
    // mocking ScanAll func of the pgxscan
    t.Run("EXPECT FAIL GETS scan data error", func(t *testing.T){
        // mock inner function pgxscan.ScanAll
        scanAll := scanAllFunc
        scanAllFunc = func(dst interface{}, rows pgx.Rows) error {
            return E.New(E.ErrDatabase)
        }
        defer func() {
            scanAllFunc = scanAll
        }()

        // prepare mock query
        mock.ExpectQuery(regexp.QuoteMeta(sqlUserStatusR)).
            WillReturnRows(pgxmock.NewRows(userStatusHeader).
                AddRow(us[0].ID, us[0].StatusName,us[0].Description).
                AddRow(us[1].ID, us[1].StatusName,us[1].Description).
                AddRow(us[2].ID, us[2].StatusName,us[2].Description),
            )

        // actual method test
        store := NewUserStatusStore(mock)
        got, err := store.Gets()
        assert.Error(t, err)
        assert.Nil(t, got)
    })
}
