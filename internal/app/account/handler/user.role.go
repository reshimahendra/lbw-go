/* 
    package handler
    user.role.go
    - handler/ interaction layer for user.role
    - NOTE of method:
    - -- UserRoleCreateHandler : method to create new user.role record
    - -- UserRoleGetHandler    : method to get user.role record by id
    - -- UserRoleGetsHandler   : method to get all user.role record
    - -- UserRoleUpdateHandler : method to update user.role record
    - -- UserRoleDeletesHandler: method to soft delete user.role record
*/
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/lbw-go/internal/app/account/service"
	"github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/helper"
)

// UserRoleHandler is type wrapper for user.role service interface
type UserRoleHandler struct {
    // Service is interfaces to user.role service
    Service service.IUserRoleService
}

// NewUserRoleHandler is instance to UserRoleHandler
func NewUserRoleHandler(Service service.IUserRoleService) *UserRoleHandler{
    return &UserRoleHandler{Service}
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
    response, err := h.Service.Create(*uReq)
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

// UserRoleGetHandler is handler to get user.role record based on its id
func (h *UserRoleHandler) UserRoleGetHandler(c *gin.Context) {
    // get 'id' param from the request context
    paramId := c.Param("id")
    id, err := strconv.Atoi(paramId)
    if err != nil {
        e := E.New(E.ErrParamIsInvalid)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // send request to service layer to retreive user.role record
    response, err := h.Service.Get(id)
    if err != nil {
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // send response to client
    helper.APIResponse(
        c,
        http.StatusOK,
        "success get user.role data",
        response,
    )
}

// UserRoleGetsHandler is handler to get all user.role record
func (h *UserRoleHandler) UserRoleGetsHandler(c *gin.Context) {
    // send request to service layer to retreive user.role record
    response, err := h.Service.Gets()
    if err != nil {
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // send response to client
    helper.APIResponse(
        c,
        http.StatusOK,
        "success get user.role data",
        response,
    )
}

// UserRoleUpdateHandler is handler to update user.role record based on its id
func (h *UserRoleHandler) UserRoleUpdateHandler(c *gin.Context) {
    // get 'id' param from the request context
    paramId := c.Param("id")
    id, err := strconv.Atoi(paramId)
    if err != nil {
        e := E.New(E.ErrParamIsInvalid)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // get new user.role data from request context
    var uReq = new(domain.UserRoleRequest)
    if err := c.ShouldBindJSON(&uReq); err != nil {
        e := E.New(E.ErrRequestDataInvalid)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }
    

    // send request to service layer to update user.role record
    response, err := h.Service.Update(id, *uReq)
    if err != nil {
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // send response to client
    helper.APIResponse(
        c,
        http.StatusOK,
        "success update user.role data",
        response,
    )
}

// UserRoleDeletesHandler is handler soft delete user.role record
func (h *UserRoleHandler) UserRoleDeletesHandler(c *gin.Context) {
    // get 'id' param from the request context
    paramId := c.Param("id")
    id, err := strconv.Atoi(paramId)
    if err != nil {
        e := E.New(E.ErrParamIsInvalid)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // send request to service layer to delete user.role record
    response, err := h.Service.Delete(id)
    if err != nil {
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // send response to client
    helper.APIResponse(
        c,
        http.StatusOK,
        "success delete user.role data",
        response,
    )
}
