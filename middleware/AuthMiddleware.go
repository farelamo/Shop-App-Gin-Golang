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
			c.JSON(http.StatusUnauthorized, gin.H{
				"Status": false,
				"message": "Invalid Token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}