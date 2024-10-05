package routes

import (
	"gfreecs0510/events/src/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoute(server *gin.Engine) {
	server.POST("/signup", controllers.SignUp)
	server.POST("/login", controllers.Login)
}
