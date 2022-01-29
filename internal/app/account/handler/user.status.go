/*
   package handler
   user.status.go
   - handler/ interaction layer for user.status
   - NOTE of method:
   - -- UserStatusGetHandler  : method to get user.status record by id
   - -- UserStatusGetsHandler : method to get all user.status record
*/
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/lbw-go/internal/app/account/service"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/helper"
)

// UserStatusHandler is type wrapper for user.status service interface
type UserStatusHandler struct {
    // Service is interfaces to user.status service
    Service service.IUserStatusService
}

// NewUserStatusHandler is instance to UserStatusHandler
func NewUserStatusHandler(Service service.IUserStatusService) *UserStatusHandler{
    return &UserStatusHandler{Service}
}

// UserStatusGetHandler is handler to get user.status record based on its id
func (h *UserStatusHandler) UserStatusGetHandler(c *gin.Context) {
    // get id param from Context
    uid := c.Param("id")
    id, err := strconv.Atoi(uid)
    if err != nil {
        e := E.New(E.ErrParamIsInvalid)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // send request to service layer to retreive user.status record
    response, err := h.Service.Get(id)
    if err != nil {
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // send response to client
    helper.APIResponse(
        c,
        http.StatusOK,
        "success get user.status data",
        response,
    )
}

// UserStatusGetsHandler is handler to get all user.status record
func (h *UserStatusHandler) UserStatusGetsHandler(c *gin.Context) {
    // send request to service layer to retreive all user.status record
    response, err := h.Service.Gets()
    if err != nil {
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // send response to client
    helper.APIResponse(
        c,
        http.StatusOK,
        "success get user.status data",
        response,
    )
}
