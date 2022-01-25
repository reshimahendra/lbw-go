package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func init() {
//     err := config.Setup()
//     if err != nil {
//         log.Println("Error on initializing configuration")
//     }
// }

// TestCheckSecureKeyLength to check whether secure key meet minimum length requirement
func TestCheckSecureKeyLength(t *testing.T) {
    cases := []struct{
        name string
        length int
        want bool
    }{
        {"SuccessA", 16, true},
        {"SuccessB", 50, true},
        {"FailA", 10, false},
    }
    
    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T){
            got := checkSecureKeyLength(tt.length)
            switch tt.want {
            case true:
                assert.Nil(t, got)                
            case false:
                assert.Error(t, got)
            }
        })
    }
}

// TestFallbackInsecureKey is to test the fallback for insecure key input
func TestFallbackInsecureKey(t *testing.T) {
    cases := []struct{
        name string
        length int
        want bool
    }{
        {"SuccessA", 16, true},
        {"SuccessB", 50, true},
        {"FailA", 10, false},
    }
    
    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T){
            got, err := fallbackInsecureKey(tt.length)
            switch tt.want {
            case true:
                assert.Nil(t, err)
                assert.NotEqual(t, got, "")
            case false:
                assert.Error(t, err)
                assert.Equal(t, got, "")
            }
        })
    }
    
}

// TestGenerateSecureKey will test the secure key generator
func TestGenerateSecureKey(t *testing.T) {
    cases := []struct{
        name string
        length int
        wantError bool
    }{
        {"EXPECT SUCCESS", 16, false},
        {"EXPECT FAIL key length less than required lenth", 15, true},
    }

    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T){
            got, err := GenerateSecureKey(tt.length)
            switch tt.wantError {
            case true:
                assert.Error(t, err)
                assert.Equal(t, got, "")
            case false:
                assert.NoError(t, err)
                assert.NotNil(t, got)
                assert.NotEqual(t, got, "")
            }
        })
    }
} 

// TestHashPassword is for testing the hash password generator
func TestHashPassword(t *testing.T) {
    cases := []struct{
        name,password string
    }{
        {"EXPECT SUCCESS alpha numeric", "1234abcd"},
        {"EXPECT SUCCESS mix char", "Aj6@1_=8"},
    }

    if testing.Short() {
        for _, tt := range cases {
            got, err:= HashPassword(tt.password)
            
            // run the test table
            t.Run(tt.name, func(t *testing.T){
                assert.NoError(t, err)
                assert.NotEmpty(t, got)                
            })
        }
   }
}

// TestCheckPasswordHash to test whether hash pass match with the plain 
// pass before it hashed
func TestCheckPasswordHash(t *testing.T) {
    cases := []struct{
        name, pass, hashPass string
        want bool
    }{
        {"SuccessA", "1234abcd", "$2a$14$./GeWe7L5mWgMxg6aa2owurPQQogskmMbpE8o2omFW4/q.W4fOtwe", true},
        {"SuccessB", "Aj6@1_=8", "$2a$14$CRqt9LW2L.ir9dyMC6G7iOw2ilwzk.r.nkj2PEu1T/j37xBX7xJHO", true},
        {"FailA", "nopass", "should-be-not-matching", false},
    }

    if testing.Short() {
        for _, tt := range cases {
            t.Run(tt.name, func(t *testing.T){
                got := CheckPasswordHash(tt.pass, tt.hashPass)
                assert.Equal(t, got, tt.want)
            })
        }
    }
}

// TestPasswordTooShort is testing whether password is too short
func TestPasswordTooShort(t *testing.T) {
    cases := []struct{
        name, pass string
        want bool
    }{
        {"SuccessA", "sweethome", false},
        {"SuccessB", "aBc%58kjh1", false},
        {"FailA", "fail", true},
        {"FailB", "", true},
    }

    for _, tt := range cases{
        t.Run(tt.name, func(t *testing.T){
            got := PasswordTooShort(tt.pass)
            assert.Equal(t, got, tt.want)
        })
    }
}

// TestEmailIsValid is for testing the validity of inputed mail
func TestEmailIsValid(t *testing.T) {
    cases := []struct{
        name, email string
        want bool
    }{
        {"SuccessA", "abc@gmail.com", true},
        {"SuccessB", "abc.cde@gmail.co.id", true},
        {"FailA", "abc.com", false},
        {"FailB", "@gmail.com", false},
    }

    for _, tt := range cases{
        t.Run(tt.name, func(t *testing.T){
            got := EmailIsValid(tt.email)
            assert.Equal(t, got, tt.want)
        })
    }
}
