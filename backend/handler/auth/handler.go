package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dia-manager-backend/service/auth"
)

func Register(c *gin.Context) {
    var req RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := auth.Register(req.Username, req.Password)

    if err != nil {
		println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}