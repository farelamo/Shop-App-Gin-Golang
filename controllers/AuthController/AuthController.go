package AuthController

import "github.com/gin-gonic/gin"

type AuthController interface {
	LoginCheck(ctx *gin.Context)
}