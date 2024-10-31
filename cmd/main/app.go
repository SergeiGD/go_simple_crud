package main

import (
	"simple_rest_crud/internal/user"
	"simple_rest_crud/internal/user/delivery/http"

	"github.com/gin-gonic/gin"
)

func main() {
	userHandlers := http.NewUserHandler(user.GetUserService())
	r := gin.Default()
	r.GET("/root/", userHandlers.HelloWorld)
	r.GET("/user/:id", userHandlers.GetUserByID)
	r.POST("/user", userHandlers.CreateUser)
	r.Run(":8080")
}
