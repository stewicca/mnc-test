package routes

import (
	"github.com/gin-gonic/gin"
	"mnc-test/controllers"
	"mnc-test/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/login", controllers.Login())
	incomingRoutes.POST("/logout", middlewares.AuthMiddleware(), controllers.Logout())
}
