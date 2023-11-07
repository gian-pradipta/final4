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
)

func main() {
	router := gin.Default()
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUser(userRepo)
	userController := controller.NewUserController(userService)

	router.POST("users/register", userController.Create)
	router.POST("users/login", userController.Login)
	router.GET("users/see", middleware.Authenticate(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"err": "Hello, world",
		})
	})

	router.Run(":8000")

}
