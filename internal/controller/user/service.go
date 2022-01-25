package user

import (
	"context"

	"github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/interfaces"
)

type UserService struct {
    Store interfaces.IUser
}

func NewUserService(rs interfaces.IUser) *UserService{
    return &UserService{Store: rs}
}

func (s *UserService) Get(ctx context.Context, id int) (*domain.UserResponse, error) {
    user, err := s.Store.Get(id)
    if err != nil {
        return nil, err
    }
    
    
    return UserToResponse(*user), nil 
}

func UserToResponse(u domain.User) *domain.UserResponse {
    // role, err := New
   return &domain.UserResponse{}
}
