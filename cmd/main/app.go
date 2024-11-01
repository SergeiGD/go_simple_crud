package main

import (
	"simple_rest_crud/internal/user/delivery/http"
	"simple_rest_crud/pkg/logging"

	_ "simple_rest_crud/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Simple CRUD docs
// @version         1.0

// @BasePath  /api/v1
func main() {
	logger := logging.GetLogger()
	logger.Info("starting app")

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		http.NewUserRoutes(v1, logger)
	}

	r.Run(":8080")
}
