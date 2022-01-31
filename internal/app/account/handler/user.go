/*
   package handler
   user.go
   - handler/ interaction layer for user.role
   - NOTE of method:
   - -- UserCreateHandler : method to create new user record
   - -- UserGetHandler    : method to get user record by id
   - -- UserGetsHandler   : method to get all user record
   - -- UserUpdateHandler : method to update user record
   - -- UserDeletesHandler: method to soft delete.role record
*/
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/lbw-go/internal/app/account/service"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/helper"
	"github.com/reshimahendra/lbw-go/internal/pkg/logger"
)

// UserHandler is type wrapper for user service interface
type UserHandler struct {
    Service service.IUserService
}

// NewUserHandler is new instance of UserHandler
func NewUserHandler(Service service.IUserService) *UserHandler{
    return &UserHandler{Service}
}

// UserCreateHandler is handler layer for Create user 
func (h *UserHandler) UserCreateHandler(c *gin.Context) {
    // get user request data from context
    req := new(d.UserRequest)
    if err := c.ShouldBindJSON(&req); err != nil {
        e := E.New(E.ErrRequestDataInvalid)
        logger.Errorf("fail binding user data: %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // send request to service layer to process insert new user record
    response, err := h.Service.Create(*req)
    if err != nil {
        logger.Errorf("fail inserting user data: %v", err)
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // send response data to user/ client
    helper.APIResponse(
        c,
        http.StatusOK,
        "success inserting user data",
        response,
    )
}

// UserGetHandler is handler layer to get user record 
func (h *UserHandler) UserGetHandler(c *gin.Context) {
    // get 'id' param from the request context
    id := c.Param("id")

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
        "success getting user data",
        response,
    )
}

// UserGetsHandler is handler layer to get all user record 
func (h *UserHandler) UserGetsHandler(c *gin.Context) {
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
        "success getting user data",
        response,
    )
}

// UserUpdateHandler is handler layer to update user 
func (h *UserHandler) UserUpdateHandler(c *gin.Context) {
    // get 'id' param from the request context
    id := c.Param("id")

    // get user request data from context
    req := new(d.UserRequest)
    if err := c.ShouldBindJSON(&req); err != nil {
        e := E.New(E.ErrRequestDataInvalid)
        logger.Errorf("fail binding user data: %v", err)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)
        return
    }

    // send request to service layer to process insert new user record
    response, err := h.Service.Update(id, *req)
    if err != nil {
        logger.Errorf("fail updating user data: %v", err)
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // send response data to user/ client
    helper.APIResponse(
        c,
        http.StatusOK,
        "success updating user data",
        response,
    )
}

// UserDeleteHandler is handler layer to delete user record 
func (h *UserHandler) UserDeleteHandler(c *gin.Context) {
    // get 'id' param from the request context
    id := c.Param("id")

    // send request to service layer to delete user record
    response, err := h.Service.Delete(id)
    if err != nil {
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // send response to client
    helper.APIResponse(
        c,
        http.StatusOK,
        "success deleting user data",
        response,
    )
}
