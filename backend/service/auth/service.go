package auth

import (
	"context"
	"errors"
	"time"

	"dia-manager-backend/config"
	"dia-manager-backend/utils"

	"github.com/golang-jwt/jwt/v5"
)

func CreateUser(username string, password string) (string, error) {

    var id string

    hashedPassword, err := utils.HashPassword(password)

    if err != nil {
        println(err)
        return "", errors.New("failed to hash password")
    }
    
    err = config.DB.QueryRow(context.Background(), `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`, username, hashedPassword).Scan(&id)

    if err != nil {
        println(err.Error())
        return "", errors.New("failed to insert the new user into the database")
    }

    return id, nil
}

func CreateToken(id string, username string) (string, error) {

    claims := jwt.MapClaims{
        "user_id":  id,
        "username": username,
        "exp":     time.Now().Add(time.Duration(config.Load().TokenLifetime) * time.Second).Unix(), // DONT LOAD CONFIG HERE
        "iat":      time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    secretKey := []byte(config.Load().TokenSecret) // DONT LOAD CONFIG HERE
    tokenString, err := token.SignedString(secretKey)
    
    if err != nil {
        println(err.Error())
        return "", errors.New("failed to create JWT token")
    }

    return tokenString, nil
}