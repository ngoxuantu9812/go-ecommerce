package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-ecomm/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.HTTPStatusCodeError(c, response.ErrTokenInvalid)
			c.Abort()
			return
		}
		c.Next()
	}
}
