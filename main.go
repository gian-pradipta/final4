package main

import (
	"final2/internal/controller"
	"final2/internal/database"
	"final2/internal/middleware"
	"final2/internal/repository"
	"final2/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUser(userRepo)
	v := validator.New()
	userController := controller.NewUserController(userService, v)

	categoryRepo := repository.NewCategory(db)
	categoryService := service.NewCategory(categoryRepo)
	categoryController := controller.NewCategory(categoryService, v)

	productRepo := repository.NewProduct(db)
	productService := service.NewProduct(productRepo)
	productController := controller.NewProduct(productService, v)
	// USER
	router.POST("users/register", userController.Create)
	router.POST("users/login", userController.Login)
	router.PATCH("users/topup", middleware.Authenticate(), userController.TopUp)
	router.GET("users/login", middleware.Authenticate(), func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusAccepted)
	})
	// CATEGORY
	router.POST("categories", middleware.Authenticate(), middleware.AuthorizeAdmin(), categoryController.Create)
	router.GET("categories", middleware.Authenticate(), categoryController.GetAll)
	router.PATCH("categories/:id", middleware.Authenticate(), middleware.AuthorizeAdmin(), categoryController.Update)
	router.DELETE("categories/:id", middleware.Authenticate(), middleware.AuthorizeAdmin(), categoryController.Delete)

	//PRODUCT
	router.POST("products", middleware.Authenticate(), middleware.AuthorizeAdmin(), productController.Create)
	router.GET("products", middleware.Authenticate(), productController.GetAll)
	router.PUT("products/:id", middleware.Authenticate(), middleware.AuthorizeAdmin(), productController.Update)
	router.DELETE("products/:id", middleware.Authenticate(), middleware.AuthorizeAdmin(), productController.Delete)
	router.Run(":8000")

}
