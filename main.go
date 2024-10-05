package main

import (
	"gfreecs0510/events/src/clients"
	"gfreecs0510/events/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Env not loaded")
	}

	clients.InitDB()
	server := gin.Default()
	routes.RegisterEventRoutes(server)
	routes.RegisterUserRoute(server)
	routes.RegisterRegistrationRoutes(server)
	server.Run(":8080")
}
