package main

import (
	"log"
	"recu_c1/Users/infraestructure"
)

func main() {
	r := infraestructure.SetupRoutes()

	log.Println("Servidor corriendo en :8080")

	r.Run(":8080")
}
