package user

import (
	"context"

	"github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/interfaces/datastore/user"
)

type UserRoleService interface {
    Get(ctx context.Context, id int) (*domain.UserRole, error)
}

type Service struct {
    RoleStore user.Datastore
}

func NewService(ds user.Datastore) Service{
    return Service{RoleStore: ds}
}

func (svc Service) Get(ctx context.Context, id int) (*domain.UserRole, error) {
    return svc.RoleStore.Get(ctx, id)
}
