package infraestructure

import (
	"net/http"
	"recu_c1/Users/application"
	"recu_c1/Users/domain"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func LongPollingHandler(service *application.UserService)  gin.HandlerFunc{
	var mu sync.Mutex
	var clients [] chan []domain.User

	return func (c *gin.Context){
		client := make(chan []domain.User, 1 )
		mu.Lock()
		clients = append(clients, client)
		mu.Unlock()

		select {
		case users := <-client:
			c.JSON(http.StatusOK, users)
		case <-time.After(1 * time.Second):
			c.JSON(http.StatusNoContent, nil) 
		}

		mu.Lock()
		for i, ch := range clients {
			if ch == client {
				clients = append(clients[:i], clients[i+1:]...)
				break
			}
		}
		mu.Unlock()
	}
}

func NotifyLongPollers(service *application.UserService, clients []chan []domain.User) {
	for _, client := range clients {
		client <- service.GetUser()
	}
}