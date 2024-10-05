package routes

import (
	"gfreecs0510/events/src/controllers"
	"gfreecs0510/events/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRegistrationRoutes(server *gin.Engine) {
	authGroup := server.Group("/")
	authGroup.Use(middleware.Authenticate)
	authGroup.POST("/events/:eventId/registration", controllers.CreateRegistration)
	authGroup.DELETE("/events/:eventId/registration", controllers.DeleteRegistration)
}
