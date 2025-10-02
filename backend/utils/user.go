package utils

import (
	"github.com/gin-gonic/gin"
)

func GetUserIdByContext(c *gin.Context) (string) {
	uidAny, _ := c.Get("userId")
	return uidAny.(string)
}

func GetUsernameByContext(c *gin.Context) (string) {
	unameAny, _ := c.Get("username")
	return unameAny.(string)
}