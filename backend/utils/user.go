package utils

import (
	"github.com/gin-gonic/gin"
)

func GetUserIdByContext(c *gin.Context) (string) {
	uidAny, _ := c.Get("userId")
	return uidAny.(string)
}