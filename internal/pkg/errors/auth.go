/*
   Package errors for authentication 
   containing custom error for our app
*/
package errors

const (
    // ErrSignUp is error code for failing to signup
    // msg = "signup fail"
    ErrSignUp = iota + 700 

    // ErrSignIn is error code for failing to signin 
    // msg = "email and password does not match or not exist"
    ErrSignIn

    // ErrSignOut is error code for failing to signout
    // msg = "signout fail"
    ErrSignOut

    // ErrUserNotRegistered is error code for unregistered access
    // msg = "user not registered"
    ErrUserNotRegistered

    // ErrUserAlreadyRegistered is error code for already registered user that tried to 
    // create new account with similar data
    // msg = "user already registered"
    ErrUserAlreadyRegistered

    // ErrUserNotActive is error code for inactive user access
    // msg = "user is not active"
    ErrUserNotActive

    // ErrPasswordNotMatch is error code for password that are not match when compared
    // msg = "password does not match"
    ErrPasswordNotMatch

    // ErrPasswordTooShort is error code for too short password
    // msg = "password is too short"
    ErrPasswordTooShort

    // ErrSecureKeyTooShort is error code for too short security key
    // msg = "secure key length does not meet minimum requirement"
    ErrSecureKeyTooShort

    // ErrTokenCreate is error code for failing creating auth token
    // msg = "could not create token"
    ErrTokenCreate

    // ErrTokenRefresh is error code for failing to refresh auth token
    // msg = "could not refresh token"
    ErrTokenRefresh

    // ErrTokenInvalid is error code for invalid auth token 
    // msg = "token invalid"
    ErrTokenInvalid

    // ErrTokenNotFound is error code for no token found
    // msg = "token not found"
    ErrTokenNotFound
)

const (
    // ErrSignUpMsg is error message for signup error
    // msg = "signup fail"
    ErrSignUpMsg = "signup fail"

    // ErrSignInMsg is error message for failing signin
    // msg = "email and password does not match or not exist"
    ErrSignInMsg = "email and password does not match or not exist"

    // ErrSignOutMsg is error message for failing to signout
    // msg = "signout fail"
    ErrSignOutMsg = "signout fail"

    // ErrUserNotRegisteredMsg is error message for unregistered access
    // msg = "user not registered"
    ErrUserNotRegisteredMsg = "user not registered"

    // ErrUserAlreadyRegisteredMsg is error message for already registered user that tried to 
    // create new account with similar data
    // msg = "user already registered"
    ErrUserAlreadyRegisteredMsg = "user already registered"

    // ErrUserNotActiveMsg is error message for inactive user access
    // msg = "user is not active"
    ErrUserNotActiveMsg = "user is not active"

    // ErrPasswordNotMatchMsg is error message for password that are not match when compared
    // msg = "password does not match"
    ErrPasswordNotMatchMsg = "password does not match"

    // ErrPasswordTooShortMsg is error message for too short password
    // msg = "password is too short"
    ErrPasswordTooShortMsg = "password is too short"

    // ErrSecureKeyTooShortMsg is error message for too short security key
    // msg = "secure key length does not meet minimum requirement"
    ErrSecureKeyTooShortMsg = "secure key length does not meet minimum requirement"

    // ErrTokenCreateMsg is error message for failing creating auth token
    // msg = "could not create token"
    ErrTokenCreateMsg = "could not create token"

    // ErrTokenRefreshMsg is error message for failing to refresh auth token
    // msg = "could not refresh token"
    ErrTokenRefreshMsg = "could not refresh token"

    // ErrTokenInvalidMsg is error message for invalid auth token 
    // msg = "token invalid"
    ErrTokenInvalidMsg = "token invalid"

    // ErrTokenNotFoundMsg is error code for no token found
    // msg = "token not found"
    ErrTokenNotFoundMsg = "token not found"
)
