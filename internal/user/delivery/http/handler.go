package http

import (
	"net/http"
	"simple_rest_crud/internal/user/domain/entity"
	"simple_rest_crud/internal/user/domain/service"
	"simple_rest_crud/pkg/logging"

	"github.com/gin-gonic/gin"
)

type userHandlers struct {
	service *service.UserService
	logger  *logging.Logger
}

func NewUserHandler(service *service.UserService, logger *logging.Logger) *userHandlers {
	return &userHandlers{service, logger}
}

func (h *userHandlers) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

// GetUserByID godoc
// @Summary      Get user by id
// @Description  Get user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  entity.UserDetailEntity
// @Router       /users/{id} 	[get]
func (h *userHandlers) GetUserByID(c *gin.Context) {
	var uri GetUserByIDDOT
	if err := c.ShouldBindUri(&uri); err != nil {
		h.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	user, err := h.service.GetUserByID(c, uri.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Error on getting user"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Not found"})
		return
	}

	c.JSON(http.StatusOK, user)

}

// CreateUser godoc
// @Summary			Create user
// @Description		Create user
// @Tags         	users
// @Accept       	json
// @Produce      	json
// @Param   		createUserRequest	body	entity.UserDetailEntity	true	"username email password"
// @Success      	200 				{integer}	int	"user id"
// @Router       	/users/				[post]
func (h *userHandlers) CreateUser(c *gin.Context) {
	var userData CreateUserDOT
	if err := c.ShouldBindJSON(&userData); err != nil {
		h.logger.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	userId, err := h.service.CreateUser(
		c,
		entity.UserCreateRawEntity{
			Email:       userData.Email,
			Username:    userData.Username,
			RawPassword: userData.Password,
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Error on creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userId})
}
