package ProductController

import (
	"fmt"
	"net/http"
	"shop/models"
	. "shop/services/ProductService"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService ProductService
}

func NewProductController(ProductService ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: ProductService,
	}
}

func (p *ProductControllerImpl) Save(ctx *gin.Context){
	var Product models.AddProduct

	if err := ctx.ShouldBindJSON(&Product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}
	
	product, err := p.ProductService.Save(&Product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Status" : true,
		"data"	 : product,
	})
}

func (p *ProductControllerImpl) FindAll(ctx *gin.Context) {
	products, err := p.ProductService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"data"	 : products,
	})
}

func (p *ProductControllerImpl) FindByCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	products, err := p.ProductService.FindByCategory(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"data"	 : products,
	})
}

func (p *ProductControllerImpl) FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	Product, err := p.ProductService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Status" : false,
			"message": "data with id " + strconv.Itoa(id) + " not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"data" 	 : Product,
	})
}

func (p *ProductControllerImpl) Update(ctx *gin.Context) {
	var Product models.Product

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&Product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status"  : false,
			"message" : err,
		})
		return
	}

	count, err := p.ProductService.Update(id, &Product)
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

func (p *ProductControllerImpl) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	count, err := p.ProductService.Delete(id)
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