package user

import (
	"dia-manager-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
    c.JSON(http.StatusCreated, gin.H{"id": utils.GetUserIdByContext(c), "username": utils.GetUsernameByContext(c)})
}