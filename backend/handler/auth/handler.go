package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"dia-manager-backend/config"
	"dia-manager-backend/service/auth"
	"dia-manager-backend/utils"
)

func Register(c *gin.Context) {
    var req RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id, err := auth.CreateUser(req.Username, req.Password)

    if err != nil {
		println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
        return
    }

    token, err := auth.CreateToken(id)

    if err != nil {
        println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(c *gin.Context) {
    var req RegisterRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var id string
    var hashedPassword string

    err := config.DB.QueryRow(context.Background(), `SELECT id, password FROM users WHERE username=$1`, req.Username).Scan(&id, &hashedPassword)

    if err != nil {
        println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user from database"})
        return
    }

    if !utils.CheckPasswordHash(req.Password, hashedPassword) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
        return
    }

    token, err := auth.CreateToken(id)

    if err != nil {
        println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}