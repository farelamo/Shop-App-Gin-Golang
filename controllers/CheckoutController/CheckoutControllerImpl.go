package CheckoutController

import (
	"net/http"
	"shop/models"
	. "shop/services/CheckoutService"
	"shop/utils/token"

	"github.com/gin-gonic/gin"
)

type CheckoutControllerImpl struct {
	CheckoutService CheckoutService
}

func NewCheckoutController(CheckoutService CheckoutService) CheckoutController {
	return &CheckoutControllerImpl{
		CheckoutService: CheckoutService,
	}
}

func (c *CheckoutControllerImpl) Save(ctx *gin.Context) {
	var cartIds models.CartCheckout

	if err := ctx.ShouldBindJSON(&cartIds); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}

	userId, err := token.GetUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}

	checkout, err := c.CheckoutService.Save(userId, &cartIds); 
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Status" : true,
		"message": checkout,
	})
}