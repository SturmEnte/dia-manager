package router

import (
	"dia-manager-backend/config"

	"github.com/gin-gonic/gin"

	authHandler "dia-manager-backend/handler/auth"

	"dia-manager-backend/middleware"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
    r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Check if request has authorization token if the request was not matched by the auth group
	r.Use(middleware.AuthMiddleware())


    return r
}
