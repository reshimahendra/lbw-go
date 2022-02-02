package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMessage is for testing error message
func TestMessage(t *testing.T) {
    cases := []struct{
        code uint
        want string
        wantErr bool
    }{
        {ErrDataIsEmpty, ErrDataIsEmptyMsg, false},
        {ErrDataNotFound, ErrDataNotFoundMsg, false},
        {ErrGettingData, ErrGettingDataMsg, false},
        {ErrInsertDataFail, ErrInsertDataFailMsg, false},
        {ErrInsertDataFail, "the message different", true},
        {1, ErrInsertDataFailMsg, true},
    }

    for _, tt := range cases {
        t.Run(tt.want, func(t *testing.T){
            got := Message(tt.code)
            switch tt.wantErr {
            case true: 
                assert.NotEqual(t, got, tt.want)
            case false:
                assert.Equal(t, got, tt.want)
            }
        })
    }
}

// TestError is for testing custom error
func TestError(t *testing.T) {
    cases := []struct{
        code uint
        msg string
    }{
        {ErrServer, ErrServerMsg},
        {ErrServerMode, ErrServerModeMsg},
        {ErrServerHost, ErrServerHostMsg},
        {ErrServerPort, ErrServerPortMsg},
        {ErrDatabase, ErrDatabaseMsg},
        {ErrDatabaseConfiguration, ErrDatabaseConfigurationMsg},
        {ErrDatabaseTransactionNil, ErrDatabaseTransactionNilMsg},
        {ErrDatabaseRollback, ErrDatabaseRollbackMsg},
        {ErrDatabasePoolNil, ErrDatabasePoolNilMsg},
        {ErrDataIsEmpty, ErrDataIsEmptyMsg},
        {ErrDataIsInvalid, ErrDataIsInvalidMsg},
        {ErrDataNotFound, ErrDataNotFoundMsg},
        {ErrGettingData, ErrGettingDataMsg},
        {ErrInsertDataFail, ErrInsertDataFailMsg},
        {ErrUpdateDataFail, ErrUpdateDataFailMsg},
        {ErrDeleteDataFail, ErrDeleteDataFailMsg},
        {ErrDataAlreadyExist, ErrDataAlreadyExistMsg},
        {ErrParamIsEmpty, ErrParamIsEmptyMsg},
        {ErrParamIsInvalid, ErrParamIsInvalidMsg},
        {ErrUsernameIsInvalid, ErrUsernameIsInvalidMsg},
        {ErrEmailIsInvalid, ErrEmailIsInvalidMsg},
        {ErrRequestDataInvalid, ErrRequestDataInvalidMsg},
    }

    for _, tt := range cases {
        t.Run(tt.msg, func(t *testing.T){
            got := New(tt.code)
            want := &Error{Code: tt.code, Message:tt.msg}
            assert.Error(t, got)
            assert.Equal(t, got.(*Error).Code, tt.code)
            assert.Equal(t, got.(*Error).Message, tt.msg)
            assert.Equal(t, got.Error(), want.Error())
            assert.Equal(t, got, want) 
        })
    }
}

// TestErrorExt is for testing extended custom error
func TestErrorExt(t *testing.T) {
    cases := []struct {
        code uint
        msg string
    }{
        {ErrSignUp, ErrSignUpMsg},
        {ErrSignIn, ErrSignInMsg},
        {ErrSignOut, ErrSignOutMsg},
        {ErrUserNotRegistered, ErrUserNotRegisteredMsg},
        {ErrUserAlreadyRegistered, ErrUserAlreadyRegisteredMsg},
        {ErrUserNotActive, ErrUserNotActiveMsg},
        {ErrPasswordNotMatch, ErrPasswordNotMatchMsg},
        {ErrPasswordTooShort, ErrPasswordTooShortMsg},
        {ErrTokenCreate, ErrTokenCreateMsg},
        {ErrTokenRefresh, ErrTokenRefreshMsg},
        {ErrTokenInvalid, ErrTokenInvalidMsg},
        {ErrTokenNotFound, ErrTokenNotFoundMsg},
    }

    for _, tt := range cases {
        t.Run(tt.msg, func(t *testing.T){
            err := New(tt.code)
            assert.Error(t, err)
            assert.Equal(t, err.(*Error).Code, tt.code)
            assert.Equal(t, err.(*Error).Message, tt.msg)

            eExt := NewExt(tt.code, err)
            assert.Error(t, eExt)
            assert.Equal(t, eExt.(*ErrorExt).Code, tt.code)
            assert.Equal(t, eExt.(*ErrorExt).Message, tt.msg)
            assert.Equal(t, eExt.(*ErrorExt).Err, err)
            assert.Equal(t, eExt.(*ErrorExt).Error(), fmt.Sprintf("Code: %d, Message: %s, Error Detail: %v", tt.code, tt.msg, err))
        })
    }
}

// TestValidationError will test input validation on user request data
func TestValidationError(t *testing.T) {
    // struct with validation input 
    type goods struct {
        ID          int     `json:"id" binding:"required,number"`
        Name        string  `json:"name" binding:"required"`
        Email       string  `json:"email" binding:"required,email"`
        Qty         int     `json:"quantity,default=0" binding:"number"`
        Description string  `json:"description"`
    }

} 
