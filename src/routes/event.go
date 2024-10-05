package routes

import (
	"gfreecs0510/events/src/controllers"
	"gfreecs0510/events/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEventViaId)

	authGroup := server.Group("/")
	authGroup.Use(middleware.Authenticate)
	authGroup.POST("/events", controllers.CreateEvent)
	authGroup.PUT("/events/:id", controllers.UpdateEvent)
	authGroup.DELETE("/events/:id", controllers.DeleteEvent)
}
