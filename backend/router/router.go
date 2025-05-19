package router

import (
	"dia-manager-backend/config"

	"github.com/gin-gonic/gin"

	authHandler "dia-manager-backend/handler/auth"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
    r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
	}

    return r
}
