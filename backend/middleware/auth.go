package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

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
        
		// TODO: Check if token is marked as invalid in db (logged out)

		// Parse and validate JWT
        jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
            // Make sure token method is HMAC
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }

            // Return secret key for validation
            return []byte(config.Load().TokenSecret), nil
        })

        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			println(err)
            return
        }

        // Check if token is valid and extract claims
        if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
            // Extract user information from claims
            userId, userIdExists := claims["user_id"].(string)
            username, usernameExists := claims["username"].(string)
            
            if !userIdExists || !usernameExists {
                c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
                return
            }

            // Set user information in context for use in handlers
            c.Set("userId", userId)
            c.Set("username", username)
        } else {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        c.Next()
    }
}