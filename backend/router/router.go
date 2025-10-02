package router

import (
	"github.com/gin-gonic/gin"

	authHandler "dia-manager-backend/handler/auth"
	catheterHandler "dia-manager-backend/handler/catheter"

	"dia-manager-backend/middleware"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

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

    return r
}
