package AuthController

import (
	"net/http"
	"shop/models"
	. "shop/services/AuthService"

	"github.com/gin-gonic/gin"
)

type AuthControllerImpl struct {
	service AuthService
}

func NewAuthController(service AuthService) AuthController {
	return &AuthControllerImpl {
		service: service,
	}
}

func (a *AuthControllerImpl) LoginCheck(ctx *gin.Context) {
	var Input models.Login

	if err := ctx.ShouldBindJSON(&Input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}

	token, err := a.service.LoginCheck(&Input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusCreated, gin.H{
		"Status" : true,
		"data"	 : token,
	})
}