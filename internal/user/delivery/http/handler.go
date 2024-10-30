package http

import (
	"fmt"
	"net/http"
	"simple_rest_crud/internal/user/domain/service"

	"github.com/gin-gonic/gin"
)

type userHandlers struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *userHandlers {
	return &userHandlers{service}
}

func (h *userHandlers) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func (h *userHandlers) GetUserByID(c *gin.Context) {
	var uri GetUserByIDDOT
	if err := c.ShouldBindUri(&uri); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"msg": err})
		return
	}

	user, err := h.service.GetUserByID(c, uri.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"msg": "Error on getting user"})
		return
	}

	if user == nil {
		c.JSON(404, gin.H{"msg": "Not found"})
		return
	}

	c.JSON(200, user)

}
