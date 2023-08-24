package main

import (
	"fmt"
	"shop/config"
	"shop/middleware"

	"shop/services/AuthService"
	"shop/services/CartService"
	"shop/services/CategoryService"
	"shop/services/CheckoutService"
	"shop/services/ProductService"
	"shop/services/UserService"

	"shop/controllers/AuthController"
	"shop/controllers/CartController"
	"shop/controllers/CategoryController"
	"shop/controllers/CheckoutController"
	"shop/controllers/ProductController"
	"shop/controllers/UserController"

	"github.com/gin-gonic/gin"
	_ "github.com/farelamo/Shop-App-Gin-Golang/APIdocs"
	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Shop API
// @version 1.0
// @description List API of Shop Application
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://swagger.io/support/
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html

// @host localhost:8000
// @Basepath /
// @schemes http
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

	checkoutService 	:= CheckoutService.NewCheckoutService(DB)
	checkoutController  := CheckoutController.NewCheckoutController(checkoutService)

	userService 		:= UserService.NewUserService(DB)
	userController 		:= UserController.NewUserController(userService)

	url   := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	group := router.Group("/api")

	/* Auth Route */
	// @Summary Login API for Shop Application
	// @Description Get JWT Token for login access
	// @Tags root
	// @Accept */*
	// @Produce json
	// @Success 200 {object} map[string]interface{}
	// @Router /api/login [post]
	group.POST("/login", authController.LoginCheck)
	group.POST("/register", userController.Save)
	
	group.Use(middleware.AuthMiddleware())
	
	group.POST("/checkout", checkoutController.Save)

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