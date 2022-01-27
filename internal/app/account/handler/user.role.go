package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/app/account/service"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/helper"
)

// UserRoleHandler is type wrapper for user.role service interface
type UserRoleHandler struct {
    // Handler is interfaces to user.role service
    Handler service.IUserRoleService
}

// NewUserRoleHandler is instance to UserRoleHandler
func NewUserRoleHandler(handler service.IUserRoleService) *UserRoleHandler{
    return &UserRoleHandler{Handler: handler}
}

// UserRoleCreateHandler is handler to Create new user.role record
func (h *UserRoleHandler) UserRoleCreateHandler(c *gin.Context) {
    // prepare instance to get user.role request dto from the request context
    var uReq = new(domain.UserRoleRequest)
    if err := c.ShouldBindJSON(uReq); err != nil {
        e := E.New(E.ErrDataIsInvalid)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // send request to service layer to process the inserting new user.role record
    response, err := h.Handler.Create(*uReq)
    if err != nil {
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // send response to client
    helper.APIResponse(
        c,
        http.StatusOK,
        "success inserting new user.role data",
        response,
    )
}
