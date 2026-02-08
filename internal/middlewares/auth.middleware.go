package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/huyle/go-ecommerce-backend-api/pkg/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authentication logic here
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrTokenInvalid, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
