/*
   package interfaces
   - 'user.role' test unit
*/
package user

import (
	"testing"
	"time"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/golang/mock/gomock"
	// "github.com/reshimahendra/lbw-go/internal/config"
	// "github.com/reshimahendra/lbw-go/internal/interfaces/datastore"
)

func TestGet(t *testing.T) {
    t.Parallel()
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
    columns := []string{"id","role_name","description","created_at","updated_at","deleted_at"}
    pgxrow := pgxpoolmock.NewRows(columns).
        AddRow(1, "test role", "test role desc", time.Now(), time.Now(), time.Now()).
        ToPgxRows()
    mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxrow, nil)

    // ctx := context.Background()
    // pool, _, err := datastore.NewDBPool(ctx, config.Get().Database)
    // if err != nil {
    //     t.Fatalf("unexpected error: %v\n", err)
    // }
    // ds := datastore.NewDatastore(pool)

    // if ds == nil {
    //     t.Fatalf("unexpected error ocur: %v\n", err)
    // }

    // t.Logf("GOT: %v\n", ds)
}
