/*
   Authentication routine with jwt
*/
package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/reshimahendra/lbw-go/internal/config"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/helper"
	"github.com/reshimahendra/lbw-go/internal/pkg/logger"
)
var (
    // generateSecureKey is instance func of helper.GenerateSecureKey
    generateSecureKeyFunc = helper.GenerateSecureKey
)

// CreateToken will 'create' a jwt token
func CreateToken(email string) (*d.TokenDetailsDTO, error) {
    // check email validity
    if !helper.EmailIsValid(email) {
        e := E.New(E.ErrEmailIsInvalid)
        return nil, e
    }

    // load server configuration
    config := config.Get()

    // prepare token detail instance
    var (
        err error
        token = new(d.TokenDetailsDTO)
    )

    token.AtExpiresTime = time.Now().Add(
        time.Duration(config.Server.AccessTokenExpireDuration) * time.Hour)
    token.RtExpiresTime = time.Now().Add(
        time.Duration(config.Server.RefreshTokenExpireDuration) * time.Hour)

    // Construct token
    atClaims := jwt.MapClaims{}
    atClaims["email"]     = email
    atClaims["user_uuid"] = "user_uuid"
    atClaims["exp"]       = time.Now().Add(time.Hour * 48).Unix()

    aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
    
    token.AccessToken, err = aToken.SignedString([]byte(config.Server.SecureKey))
    if err != nil {
        logger.Errorf("error occur while creating access token: %v\n", err)
        return nil, err
    }

    // Construct refresh token 
    rtClaims := jwt.MapClaims{}
    rtClaims["email"]     = email
    rtClaims["user_uuid"] = "user_uuid"
    rtClaims["exp"]       = time.Now().Add(time.Hour * 96).Unix()

    rToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

    // rt := NewJwtToken(rToken)
    token.RefreshToken, err = rToken.SignedString([]byte(config.Server.SecureKey))
    if err != nil {
        logger.Errorf("error occur while creating refresh token: %v\n", err)
        return nil, err
    }

    // Generate secure key
    generateKey, err := generateSecureKeyFunc(config.Server.MinimumSecureKeyLength)
    if err != nil {
        logger.Errorf("error occur while generating secure key: %v\n", err)
        return nil, err
    }
    token.TransmissionKey = generateKey

    return token, err
}

// verifyToken will verify the given token 
func verifyToken(token string) (*jwt.Token, error) {
    config := config.Get()
    verifiedToken, err := jwt.Parse(token, func (verifiedToken *jwt.Token) (interface{}, error) {
        if _, ok := verifiedToken.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", verifiedToken.Header["alg"])
        }
        return []byte(config.Server.SecureKey), nil
    })

    if err != nil {
        e := E.New(E.ErrTokenInvalid)
        return verifiedToken, e 
    }

    return verifiedToken, nil
}

// TokenValid will check whether the 'given' token was valid or not
func TokenValid(bearerToken string) (*jwt.Token, error) {
    // new invalid token error
    e := E.New(E.ErrTokenInvalid)
    
    // check token validity
    token, err := verifyToken(bearerToken)
    if err != nil {
        if token != nil {
            return token, e
        }
        return nil, e
    }

    if !token.Valid {
        return nil, e
    }

    return token, nil
}
