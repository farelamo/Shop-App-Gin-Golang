package CategoryController

import (
	"fmt"
	"net/http"
	"shop/models"
	. "shop/services/CategoryService"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	CategoryService CategoryService
}

func NewCategoryController(CategoryService CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: CategoryService,
	}
}

func (c *CategoryControllerImpl) Save(ctx *gin.Context){
	var Category models.AddCategory

	if err := ctx.ShouldBindJSON(&Category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}
	
	category, err := c.CategoryService.Save(&Category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Status" : true,
		"data"	 : category,
	})
}

func (c *CategoryControllerImpl) FindAll(ctx *gin.Context) {
	categories, err := c.CategoryService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"data"	 : categories,
	})
}

func (c *CategoryControllerImpl) FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	category, err := c.CategoryService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Status" : false,
			"message": "data with id " + strconv.Itoa(id) + " not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"data" 	 : category,
	})
}

func (c *CategoryControllerImpl) Update(ctx *gin.Context) {
	var category models.Category

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status"  : false,
			"message" : err,
		})
		return
	}

	count, err := c.CategoryService.Update(id, &category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}

	message := fmt.Sprintf("Updated data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"message": message,
	})
}

func (c *CategoryControllerImpl) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	count, err := c.CategoryService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("Deleted data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"message": message,
	})	
}