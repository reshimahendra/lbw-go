/*
    package errors
    handler.go
    - error constant for handling handler error
*/
package errors

const (
    // ErrParamIDEmpty is error code for empty param that retreived from 'context'
    // msg = "request parameter is empty"
    ErrParamIsEmpty = iota + 600

    // ErrParamInvalid is error code for invalid request param
    // msg = "request parameter is invalid"
    ErrParamIsInvalid

    // ErrUsernameIsInvalid is error code for invalid username
    // msg = "username is invalid"
    ErrUsernameIsInvalid

    // ErrEmailIsInvalid is error code for invalid email
    // msg = "email is invalid"
    ErrEmailIsInvalid

    // ErrRequestDataInvalid is error code for invalid request data
    // msg = "request data invalid"
    ErrRequestDataInvalid
)
const (
    // ErrParamEmpty is error code for empty param that retreived from 'context'
    // msg = "request parameter is empty"
    ErrParamIsEmptyMsg = "request parameter is empty"

    // ErrParamInvalidMsg is error message for invalid request param
    // msg = "request parameter is invalid"
    ErrParamIsInvalidMsg = "request parameter is invalid"

    // ErrUsernameIsInvalidMsg is error message for invalid username
    // msg = "username is invalid"
    ErrUsernameIsInvalidMsg = "username is invalid"

    // ErrEmailIsInvalidMsg is error code for invalid email
    // msg = "email is invalid"
    ErrEmailIsInvalidMsg = "email is invalid"

    // ErrRequestDataInvalid is error message for invalid request data
    // msg = "request data invalid"
    ErrRequestDataInvalidMsg = "request data invalid"
)
