package controllers

import (
	"github.com/gin-gonic/gin"
	"mnc-test/services"
	"net/http"
)

func Transfer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			RecipientID uint `json:"recipient_id" binding:"required"`
			Amount      int  `json:"amount" binding:"required"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		senderID := c.GetUint("sub")
		if err := services.Transfer(senderID, body.RecipientID, body.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Transfer successfully"})
	}
}
