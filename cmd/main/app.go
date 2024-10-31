package main

import (
	"simple_rest_crud/internal/user"
	"simple_rest_crud/internal/user/delivery/http"
	"simple_rest_crud/pkg/logging"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("starting app")

	userHandlers := http.NewUserHandler(user.GetUserService(logger), logger)
	r := gin.Default()
	r.GET("/root/", userHandlers.HelloWorld)
	r.GET("/user/:id", userHandlers.GetUserByID)
	r.POST("/user", userHandlers.CreateUser)
	r.Run(":8080")
}
