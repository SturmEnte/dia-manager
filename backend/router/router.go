package router

import (
	"dia-manager-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
    r := gin.Default()

	r.GET("/hello", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "World"})
	})

    return r
}
