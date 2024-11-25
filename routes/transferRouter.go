package routes

import (
	"github.com/gin-gonic/gin"
	"mnc-test/controllers"
)

func TransferRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/transfer", controllers.Transfer())
}
