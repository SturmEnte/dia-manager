package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"dia-manager-backend/config"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        
        token, err := c.Cookie("token")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error while checking token validity"})
			println(err)
			return
		}
        
		var expires time.Time
		var userId string
		config.DB.QueryRow(context.Background(), "SELECT expires, user_id FROM sessions WHERE token=$1", token).Scan(&expires, &userId)

		if userId == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if expires.UnixMilli() <= time.Now().UnixMilli() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
            return
		}

		c.Set("userId", userId)

        c.Next()
    }
}