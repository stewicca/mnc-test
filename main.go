package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mnc-test/middlewares"
	"mnc-test/repositories"
	"mnc-test/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	_ = repositories.CleanTokenExpired()

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middlewares.AuthMiddleware())
	routes.TransferRoutes(router)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
