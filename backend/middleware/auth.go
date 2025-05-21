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
        
        token := c.GetHeader("Authorization")

        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        
		var expires time.Time
		config.DB.QueryRow(context.Background(), "SELECT expires FROM sessions WHERE token=$1", token).Scan(&expires)

		if expires.UnixMilli() <= time.Now().UnixMilli() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
            return
		}

        c.Next()
    }
}