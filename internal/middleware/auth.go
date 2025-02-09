// Package middleware internal/middleware/auth.go
package middleware

import (
	"github.com/gin-gonic/gin"
	"go1/pkg/e"
	"go1/pkg/response"
	"go1/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.Error(c, e.CodeUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			response.Error(c, e.CodeUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}
