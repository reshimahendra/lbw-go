/*
   Package errors for authentication 
   containing custom error for our app
*/
package errors

const (
    // ErrSignUp is error code for failing to signup
    ErrSignUp = iota + 700 

    // ErrSignIn is error code for failing to signin 
    ErrSignIn

    // ErrSignOut is error code for failing to signout
    ErrSignOut

    // ErrUserNotRegistered is error code for unregistered access
    ErrUserNotRegistered

    // ErrUserAlreadyRegistered is error code for already registered user that tried to 
    // create new account with similar data
    ErrUserAlreadyRegistered

    // ErrUserNotActive is error code for inactive user access
    ErrUserNotActive

    // ErrPasswordNotMatch is error code for password that are not match when compared
    ErrPasswordNotMatch

    // ErrPasswordTooShort is error code for too short password
    ErrPasswordTooShort

    // ErrSecureKeyTooShort is error code for too short security key
    ErrSecureKeyTooShort

    // ErrTokenCreate is error code for failing creating auth token
    ErrTokenCreate

    // ErrTokenRefresh is error code for failing to refresh auth token
    ErrTokenRefresh

    // ErrTokenInvalid is error code for invalid auth token 
    ErrTokenInvalid

    // ErrTokenNotFound is error code for no token found
    ErrTokenNotFound
)

const (
    // ErrSignUpMsg is error message for signup error
    ErrSignUpMsg = "signup fail"

    // ErrSignInMsg is error message for failing signin
    ErrSignInMsg = "email and password does not match or not exist"

    // ErrSignOutMsg is error message for failing to signout
    ErrSignOutMsg = "signout fail"

    // ErrUserNotRegisteredMsg is error message for unregistered access
    ErrUserNotRegisteredMsg = "user not registered"

    // ErrUserAlreadyRegisteredMsg is error message for already registered user that tried to 
    // create new account with similar data
    ErrUserAlreadyRegisteredMsg = "user already registered"

    // ErrUserNotActiveMsg is error message for inactive user access
    ErrUserNotActiveMsg = "user is not active"

    // ErrPasswordNotMatchMsg is error message for password that are not match when compared
    ErrPasswordNotMatchMsg = "password does not match"

    // ErrPasswordTooShortMsg is error message for too short password
    ErrPasswordTooShortMsg = "password is too short"

    // ErrSecureKeyTooShortMsg is error message for too short security key
    ErrSecureKeyTooShortMsg = "secure key length does not meet minimum requirement"

    // ErrTokenCreateMsg is error message for failing creating auth token
    ErrTokenCreateMsg = "could not create token"

    // ErrTokenRefreshMsg is error message for failing to refresh auth token
    ErrTokenRefreshMsg = "could not refresh token"

    // ErrTokenInvalidMsg is error message for invalid auth token 
    ErrTokenInvalidMsg = "token invalid"

    // ErrTokenNotFoundMsg is error code for no token found
    ErrTokenNotFoundMsg = "token not found"
)
