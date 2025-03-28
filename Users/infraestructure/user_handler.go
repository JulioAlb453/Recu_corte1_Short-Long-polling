package infraestructure

import (
	"net/http"
	"recu_c1/Users/application"
	"recu_c1/Users/domain"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *application.UserService
}

func NewUserService(service *application.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) AddUser(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Datos invalidos"})
		return
	}

	h.service.AddUser(user)
	c.Status(http.StatusCreated)
}
