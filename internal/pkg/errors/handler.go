package errors

const (
    // ErrParamIDEmpty is error code for empty param that retreived from 'context'
    ErrParamIsEmpty = iota + 600

    // ErrParamInvalid is error code for invalid request param
    ErrParamIsInvalid

    // ErrUsernameIsInvalid is error code for invalid username
    ErrUsernameIsInvalid

    // ErrEmailIsInvalid is error code for invalid email
    ErrEmailIsInvalid

    // ErrRequestDataInvalid is error code for invalid request data
    ErrRequestDataInvalid
)
const (
    // ErrParamEmpty is error code for empty param that retreived from 'context'
    ErrParamIsEmptyMsg = "request parameter is empty"

    // ErrParamInvalidMsg is error message for invalid request param
    ErrParamIsInvalidMsg = "request parameter is invalid"

    // ErrUsernameIsInvalidMsg is error message for invalid username
    ErrUsernameIsInvalidMsg = "username is invalid"

    // ErrEmailIsInvalidMsg is error code for invalid email
    ErrEmailIsInvalidMsg = "email is invalid"

    // ErrRequestDataInvalid is error code for invalid request data
    // ErrRequestDataInvalid is error code for invalid request data
    ErrRequestDataInvalidMsg = "request data invalid"
)
