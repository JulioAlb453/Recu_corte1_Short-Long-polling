package infraestructure

import (
	"recu_c1/Users/application"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	repo := NewInMemomyUserRepository()
	userService := application.NewUserService(repo)
	handler := NewUserService(userService)

	r.POST("/addPerson", handler.AddUserHandler)

	return r
}