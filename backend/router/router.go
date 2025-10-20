package router

import (
	"github.com/gin-gonic/gin"

	"dia-manager-backend/config"

	authHandler "dia-manager-backend/handler/auth"
	catheterHandler "dia-manager-backend/handler/catheter"
	userHandler "dia-manager-backend/handler/user"

	"dia-manager-backend/middleware"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
    r := gin.Default()

	// Inject config into context for all requests
    r.Use(func(c *gin.Context) {
        c.Set("config", cfg)
        c.Next()
    })

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.DELETE("/logout", authHandler.Logout)
	}

	// Check if request has authorization token if the request was not matched by the auth group
	r.Use(middleware.AuthMiddleware())

	catheter := r.Group("/catheters")
    {
        catheter.POST("", catheterHandler.CreateCatheter)
        catheter.GET("", catheterHandler.GetCatheters)
        catheter.GET("/:id", catheterHandler.GetCatheterByID)
        catheter.PUT("/:id", catheterHandler.UpdateCatheter)
        catheter.DELETE("/:id", catheterHandler.DeleteCatheter)
    }

	user := r.Group("/user")
	{
		user.GET("/me", userHandler.GetUserInfo)
	}

    return r
}
