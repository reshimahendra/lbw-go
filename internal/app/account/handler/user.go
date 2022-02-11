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
   - -- UserSignupHandler : method to signup (create new user)
   - -- UserSigninHandler : method to signin/ login
*/
package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/reshimahendra/lbw-go/internal/app/account/service"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	"github.com/reshimahendra/lbw-go/internal/pkg/auth"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/helper"
	"github.com/reshimahendra/lbw-go/internal/pkg/logger"
)

var (
    // checkPasswordHashFunc is instance func wrapper for helper.CheckPasswordHash
    checkPasswordHashFunc = helper.CheckPasswordHash

    // createToken is instance func wrapper for auth.CreateToken
    createTokenFunc = auth.CreateToken
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

// SignupHandler is handler/ controller to sign up new user
func (h *UserHandler) SignupHandler(c *gin.Context) {    
    var userRequest d.UserRequest

    err := c.ShouldBindJSON(&userRequest)
    if err != nil {
        e := E.New(E.ErrRequestDataInvalid)
        logger.Errorf("%s. %v", E.ErrSignUpMsg, err)
        helper.APIErrorResponse(c, http.StatusBadRequest, e)

        return
    }

    // check if user already exist
    isUserExist := h.Service.IsUserExist(userRequest.Username, userRequest.Email)
    if isUserExist {
        err := E.New(E.ErrUserAlreadyRegistered)
        logger.Errorf("%s. %v", E.ErrSignUpMsg, err)
        helper.APIErrorResponse(c, http.StatusBadRequest, err)

        return
    }

    // create user account. exit if error
    userResponse, err := h.Service.Create(userRequest)
    if err != nil {
        logger.Errorf("%s. %v", E.ErrSignUpMsg, err)
        helper.APIErrorResponse(c, http.StatusInternalServerError, err)

        return
    }

    // send response to client
    helper.APIResponse(
        c,
        http.StatusOK,
        "success signup",
        userResponse,
    )
}

// SigninHandler is handler/ controller to sign in/login user
func (h *UserHandler) SigninHandler(c *gin.Context) {
    // get login data from context
    var login d.AuthLoginDTO
    err := c.ShouldBindJSON(&login)
    if err != nil {
        e := E.New(E.ErrRequestDataInvalid)
        logger.Errorf("%s: %v", E.ErrRequestDataInvalidMsg, err)
        helper.APIErrorResponse(c, http.StatusUnauthorized, e)

        return
    }

    // get credential by its username
    cred, err := h.Service.GetByEmail(login.Email)
    if err != nil {
        logger.Errorf("login fail: %v", err)
        helper.APIErrorResponse(c, http.StatusUnauthorized, err)
        return
    }

    // User not active 
    if !cred.IsActive() {
        err := E.New(E.ErrUserNotActive)
        logger.Errorf("login fail: %v", err)
        helper.APIErrorResponse(c, http.StatusUnauthorized, err)

        return
    }

    // Check whether user & password match
    isPasswordMatch := checkPasswordHashFunc(login.Passkey, cred.PassKey)
    if !isPasswordMatch {
        err := E.New(E.ErrSignIn)
        logger.Errorf("login fail: %v", err)
        helper.APIErrorResponse(c, http.StatusUnauthorized, err)

        return
    }

    if cred.IsActive() && isPasswordMatch {
        token, err := auth.CreateToken(login.Email)
        if err != nil {
            e := E.New(E.ErrTokenCreate)
            logger.Errorf("%s: %v", E.ErrTokenCreateMsg, e)
            helper.APIErrorResponse(c, http.StatusInternalServerError, e)

            return
        }

        // send token data response to the client
        authLoginResponse := d.AuthLoginResponse{
            AccessToken     : token.AccessToken,
            RefreshToken    : token.RefreshToken,
            TransmissionKey : token.TransmissionKey,
        }

        helper.APIResponse(
            c,
            http.StatusOK,
            "success signin",
            authLoginResponse,
        )
    }
}

// RefreshTokenHandler is handler to refresh user account access token
func (h *UserHandler) RefreshTokenHandler(c *gin.Context) {
    mapToken := map[string]string{}

    decoder := json.NewDecoder(c.Request.Body)
    if err := decoder.Decode(&mapToken); err != nil {
        logger.Errorf("decode json fail on refresh token: %v", err)
        e := E.New(E.ErrTokenRefresh)
        helper.APIErrorResponse(c, http.StatusUnprocessableEntity, e)
        return
    }

    defer c.Request.Body.Close()

    token, err := auth.TokenValid(mapToken["refresh_token"])
    if err != nil {
        logger.Errorf("token validity check fail on refresh token: %v", err)
        e := E.New(E.ErrTokenInvalid)
        helper.APIErrorResponse(c, http.StatusUnauthorized, e)

        return
    }

    email := token.Claims.(jwt.MapClaims)["email"].(string)

    // Create new token 
    newToken, err := createTokenFunc(email)
    if err != nil {
        e := E.New(E.ErrTokenCreate)
        logger.Errorf("token creation fail on refresh token: %v", err)
        helper.APIErrorResponse(c, http.StatusInternalServerError, e)

        return
    }

    // send RefreshToken data response to the client
    authLoginResponse := d.AuthLoginResponse{
        AccessToken : newToken.AccessToken,
        RefreshToken : newToken.RefreshToken,
        TransmissionKey : newToken.TransmissionKey,
    }

    helper.APIResponse(
        c,
        http.StatusOK,
        "success refresh token",
        authLoginResponse,
    )
}

// CheckTokenHandler is handler to check user account token
func (h *UserHandler) CheckTokenHandler(c *gin.Context) {
    var decToken string
    bearerToken := c.GetHeader("Authorization")
    authArray := strings.Split(bearerToken, " ")
    if len(authArray) == 2 {
        decToken = authArray[1]
    }

    if decToken == "" {
        err := E.New(E.ErrTokenNotFound)
        logger.Errorf("check token fail: %v", err)
        helper.APIErrorResponse(c, http.StatusUnauthorized, err)

        return
    }

    token, err := auth.TokenValid(decToken)
    if err != nil {
        e := E.New(E.ErrTokenInvalid)
        logger.Errorf("token validity check fail on CheckToken: %v", err)
        helper.APIErrorResponse(c, http.StatusUnauthorized, e)

        return
    }

    // send response of the 'checkToken' result to client
    email := token.Claims.(jwt.MapClaims)["email"].(string)
    helper.APIResponse(
        c,
        http.StatusOK,
        "success checking token",
        email,
    )
}

