package main

import (
	"fmt"
	"shop/config"
	"shop/middleware"

	"shop/services/AuthService"
	"shop/services/CategoryService"
	"shop/services/ProductService"
	"shop/services/UserService"
	"shop/services/CartService"

	"shop/controllers/AuthController"
	"shop/controllers/CategoryController"
	"shop/controllers/ProductController"
	"shop/controllers/UserController"
	"shop/controllers/CartController"

	"github.com/gin-gonic/gin"
)

func main(){
	router 	:= gin.Default()
	DB 		:= config.Connect()

	authService 		:= AuthService.NewAuthService(DB)
	authController 		:= AuthController.NewAuthController(authService)

	categoryService 	:= CategoryService.NewCategoryService(DB)
	categoryController 	:= CategoryController.NewCategoryController(categoryService)

	productService 		:= ProductService.NewProductService(DB)
	productController 	:= ProductController.NewProductController(productService)

	cartService 		:= CartService.NewCartService(DB)
	cartController 		:= CartController.NewCartController(cartService)

	userService 		:= UserService.NewUserService(DB)
	userController 		:= UserController.NewUserController(userService)

	group := router.Group("/api")
	
	/* Auth Route */
	group.POST("/login", authController.LoginCheck)
	group.POST("/register", userController.Save)
	
	group.Use(middleware.AuthMiddleware())
	
	/* Cart Route */
	group.GET("/cart", cartController.FindAll)
	group.POST("/cart", cartController.Save)
	group.DELETE("/cart/:id", cartController.Delete)

	/* Category Route */
	group.GET("/category", categoryController.FindAll)
	group.GET("/category/:id", categoryController.FindById)
	group.POST("/category", categoryController.Save)
	group.PUT("/category/:id", categoryController.Update)
	group.DELETE("/category/:id", categoryController.Delete)

	/* Product Route */
	group.GET("/product", productController.FindAll)
	group.GET("/product/:id", productController.FindById)
	group.GET("/product/category/:id", productController.FindByCategory)
	group.POST("/product", productController.Save)
	group.PUT("/product/:id", productController.Update)
	group.DELETE("/product/:id", productController.Delete)

	/* User Route */
	group.GET("/user", userController.FindAll)
	group.GET("/user/:id", userController.FindById)
	group.PUT("/user/:id", userController.Update)
	group.DELETE("/user/:id", userController.Delete)

	runWithPort := fmt.Sprintf("0.0.0.0:%s", "8000")
	router.Run(runWithPort)
}