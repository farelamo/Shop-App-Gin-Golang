package CheckoutController

import (
	"github.com/gin-gonic/gin"
)

type CheckoutController interface {
	Save(ctx *gin.Context)
}