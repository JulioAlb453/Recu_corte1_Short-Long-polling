package main

import (
	"log"
	"recu_c1/Users/infraestructure"
	"github.com/gin-contrib/cors"

)

func main() {
	r := infraestructure.SetupRoutes()
	r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, 
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	log.Println("Servidor corriendo en :8080")

	r.Run(":8080")
}
