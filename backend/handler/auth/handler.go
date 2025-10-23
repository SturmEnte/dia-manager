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

    cfg := c.MustGet("config").(*config.Config)
    token, err := auth.CreateToken(cfg, id, req.Username)

    if err != nil {
        println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
        return
    }

    c.SetCookie("token", token, 3600, "/", "", false, true)
    c.Status(http.StatusOK)
}

func Login(c *gin.Context) {
    var req LoginRequest

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

    cfg := c.MustGet("config").(*config.Config)

    token, err := auth.CreateToken(cfg, id, req.Username)

    if err != nil {
        println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
        return
    }

    c.SetCookie("token", token, cfg.TokenLifetime * 60, "/", "", false, true)
    c.Status(http.StatusOK)
}

func Logout(c *gin.Context) {
    
    token, err := c.Cookie("token")

    // There should be no error at this point because the cookie was already checked by the middleware
    if err != nil {
        println(err.Error())
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    cfg := c.MustGet("config").(*config.Config)

    err = auth.DisableToken(cfg, token)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusOK)
}