package controllers

import (
	"github.com/gin-gonic/gin"
	"mnc-test/services"
	"net/http"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := services.Login(body.Username, body.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetString("token")
		username := c.GetString("username")
		exp := int64(c.GetFloat64("exp"))

		err := services.Logout(token, username, exp)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "logged out"})
	}
}
