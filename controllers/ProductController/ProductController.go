package ProductController

import "github.com/gin-gonic/gin"

type ProductController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindByCategory(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}