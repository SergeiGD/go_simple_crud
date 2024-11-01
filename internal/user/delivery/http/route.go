package http

import (
	"simple_rest_crud/internal/user"
	"simple_rest_crud/pkg/logging"

	"github.com/gin-gonic/gin"
)

type userRoutes struct {
	parent *gin.RouterGroup
	logger *logging.Logger
}

func NewUserRoutes(parent *gin.RouterGroup, logger *logging.Logger) userRoutes {
	r := userRoutes{
		parent: parent,
		logger: logger,
	}

	userHandlers := NewUserHandler(user.GetUserService(logger), logger)

	users := parent.Group("/users")
	{
		users.GET("/root/", userHandlers.HelloWorld)
		users.GET("/:id", userHandlers.GetUserByID)
		users.POST("/user", userHandlers.CreateUser)
	}

	return r
}
