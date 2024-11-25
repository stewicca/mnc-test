package middlewares

import (
	"github.com/gin-gonic/gin"
	"mnc-test/helpers"
	"mnc-test/repositories"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No Authorization header provided"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := helpers.VerifyToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
			c.Abort()
			return
		}

		blacklist, err := repositories.LoadTokenBlacklist()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error load token blacklist"})
			c.Abort()
			return
		}

		for _, blacklistedToken := range blacklist {
			if blacklistedToken.Token == tokenStr {
				if blacklistedToken.ExpiresAt.After(time.Now()) {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Please login again"})
					c.Abort()
					return
				}
			}
		}

		c.Set("token", tokenStr)
		c.Set("sub", claims["sub"])
		c.Set("username", claims["username"])
		c.Set("exp", claims["exp"])

		c.Next()
	}
}
