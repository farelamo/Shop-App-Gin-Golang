package main

import (
	"fmt"
	"shop/config"

	"shop/services/CategoryService"
	"shop/services/ProductService"
	"shop/services/UserService"

	"shop/controllers/CategoryController"
	"shop/controllers/ProductController"
	"shop/controllers/UserController"

	"github.com/gin-gonic/gin"
)

func main(){
	router 	:= gin.Default()
	DB 		:= config.Connect()

	categoryService 	:= CategoryService.NewCategoryService(DB)
	categoryController 	:= CategoryController.NewCategoryController(categoryService)

	productService 		:= ProductService.NewProductService(DB)
	productController 	:= ProductController.NewProductController(productService)

	userService 		:= UserService.NewUserService(DB)
	userController 		:= UserController.NewUserController(userService)

	/* Category Route */
	router.GET("/category", categoryController.FindAll)
	router.GET("/category/:id", categoryController.FindById)
	router.POST("/category", categoryController.Save)
	router.PUT("/category/:id", categoryController.Update)
	router.DELETE("/category/:id", categoryController.Delete)

	/* Product Route */
	router.GET("/product", productController.FindAll)
	router.GET("/product/:id", productController.FindById)
	router.GET("/product/category/:id", productController.FindByCategory)
	router.POST("/product", productController.Save)
	router.PUT("/product/:id", productController.Update)
	router.DELETE("/product/:id", productController.Delete)

	/* User Route */
	router.GET("/user", userController.FindAll)
	router.GET("/user/:id", userController.FindById)
	router.POST("/user", userController.Save)
	router.PUT("/user/:id", userController.Update)
	router.DELETE("/user/:id", userController.Delete)

	runWithPort := fmt.Sprintf("0.0.0.0:%s", "8000")
	router.Run(runWithPort)
}