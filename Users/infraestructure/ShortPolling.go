package infraestructure

import (
	"recu_c1/Users/application"

	"github.com/gin-gonic/gin"
)

func ShortPollingHandler(service *application.UserService) gin.HandlerFunc{
	return func(c *gin.Context){
		users := service.GetUser()
		c.JSON(200, gin.H{"users": users})
	}
}