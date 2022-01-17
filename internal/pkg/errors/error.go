/*
   Package errors
   * processing custom error
*/
package errors

import (
	"fmt"
)

// Error is custom error struct that only showing simple error info
type Error struct {
    Code    uint    `json:"code"`
    Message string  `json:"message"`
}

// Error method for displaying error string for 'Error' struct
func (e *Error) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// New will create new 'Error' instance 
func New(code uint) error {
    return &Error{
        Code    : code,
        Message : Message(code),
    }
}

// ErrorExt is custom error struct that can 'pass' the main Error detail via 'Err' field
type ErrorExt struct {
    Code    uint        `json:"code"`
    Message string      `json:"message,omitempty"`
    Err     interface{} `json:"error,omitempty"`
}

// Error method for displaying error string for 'Error' struct
func (e *ErrorExt) Error() string{
    return fmt.Sprintf("Code: %d, Message: %s, Error Detail: %v", e.Code, e.Message, e.Err)
}

// NewExt will create new 'ErrorExt' error instance
func NewExt(code uint, e error) error { 
    return &ErrorExt{
        Code    : uint(code),
        Message : Message(uint(code)),
        Err     : e,
    }
}

// Message wiil return error message
func Message(code uint) (message string) {
    switch code {
        // database error
        case ErrDatabase                : message = ErrDatabaseMsg 
        case ErrDatabaseTransactionNil  : message = ErrDatabaseTransactionNilMsg 
        case ErrDatabaseRollback        : message = ErrDatabaseRollbackMsg 
        case ErrDatabasePoolNil         : message = ErrDatabasePoolNilMsg 
        case ErrDataIsEmpty             : message = ErrDataIsEmptyMsg 
        case ErrDataNotFound            : message = ErrDataNotFoundMsg
        case ErrGettingData             : message = ErrGettingDataMsg
        case ErrSaveDataFail            : message = ErrSaveDataFailMsg
        case ErrUpdateDataFail          : message = ErrUpdateDataFailMsg 
        case ErrDeleteDataFail          : message = ErrDeleteDataFailMsg
        case ErrDataAlreadyExist        : message = ErrDataAlreadyExistMsg
        
        // auth error
        case ErrSignUp                  : message = ErrSignUpMsg 
        case ErrSignIn                  : message = ErrSignInMsg 
        case ErrSignOut                 : message = ErrSignOutMsg
        case ErrUserNotRegistered       : message = ErrUserNotRegisteredMsg 
        case ErrUserAlreadyRegistered   : message = ErrUserAlreadyRegisteredMsg 
        case ErrUserNotActive           : message = ErrUserNotActiveMsg 
        case ErrPasswordNotMatch        : message = ErrPasswordNotMatchMsg
        case ErrPasswordTooShort        : message = ErrPasswordTooShortMsg
        case ErrTokenCreate             : message = ErrTokenCreateMsg
        case ErrTokenRefresh            : message = ErrTokenRefreshMsg
        case ErrTokenInvalid            : message = ErrTokenInvalidMsg
        case ErrTokenNotFound           : message = ErrTokenNotFoundMsg

        // handler error
        case ErrParamIsEmpty        : message = ErrParamIsEmptyMsg
        case ErrParamIsInvalid      : message = ErrParamIsInvalidMsg
        case ErrUsernameIsInvalid   : message = ErrUsernameIsInvalidMsg
        case ErrEmailIsInvalid      : message = ErrEmailIsInvalidMsg
        case ErrRequestDataInvalid  : message = ErrRequestDataInvalidMsg

        // default/ unknown error
        default                     : message = "unknown error"
    }

    return
}

// ValidationError is a 'Request' error detail generator 
// it will break given error value into detailed error message 
// so we know which field is the cause of error upon handling request
// func ValidationError(err error) (eMessage *[]string) {
//     // Create list of error message
//     eMsg := []string{}
//     for _, e := range err.(validator.ValidationErrors) {
//         msg := fmt.Sprintf("Field error :'%s', condition: '%s'", e.Field(), e.ActualTag())
//         eMsg = append(eMsg, msg)
//     }
//
//     eMessage = &eMsg
//
//     return
// }
