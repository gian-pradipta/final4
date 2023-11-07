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

	router.POST("users/register", userController.Create)
	router.POST("users/login", userController.Login)
	router.PATCH("users/topup", middleware.Authenticate(), userController.TopUp)
	router.GET("users/login", middleware.Authenticate(), func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusAccepted)
	})

	router.Run(":8000")

}
