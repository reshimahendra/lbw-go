package handler

import "github.com/reshimahendra/lbw-go/internal/app/account/service"

type UserHandler struct {
    Handler service.IUserService
}

func NewUserHandler(handler service.IUserService) *UserHandler{
    return &UserHandler{Handler: handler}
}
