package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"shop/utils/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		err := token.TokenValidation(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unathorized")
			c.Abort()
			return
		}

		c.Next()
	}
}