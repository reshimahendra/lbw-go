/*
   Middleware to prevent unauthorized access
*/
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/lbw-go/internal/pkg/auth"
	"github.com/reshimahendra/lbw-go/internal/pkg/helper"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
)

// Authorize is middleware to prevent unauthorized access
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string
		bearerToken := c.GetHeader("Authorization")
		strArr := strings.Split(bearerToken, " ")
		if len(strArr) == 2 {
			tokenStr = strArr[1]
		}

		if tokenStr == "" {
            e := E.New(E.ErrTokenNotFound)
            helper.APIErrorResponse(c, http.StatusUnauthorized, e)
			return
		}

		token, err := auth.TokenValid(tokenStr)
		if err != nil {
            helper.APIErrorResponse(c, http.StatusUnauthorized, err)
			return
		}

		if err != nil && !token.Valid {
            helper.APIErrorResponse(c, http.StatusUnauthorized, err)
            return
		}

	}
}
