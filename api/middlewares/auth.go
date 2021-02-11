package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authorization middleware
func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := findAuthorizationToken(ctx.Request)

		// Validate token
		if tokenStr != "" {
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte("someSecret"), nil
			})

			if err == nil && token.Valid {
				ctx.Next()
			} else {
				ctx.AbortWithError(401, errors.New("unauthorized"))
				return
			}
		}

		if tokenStr == "" {
			ctx.AbortWithError(401, errors.New("unauthorized"))
			return
		}

	}
}

func findAuthorizationToken(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
