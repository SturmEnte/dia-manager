package auth

import (
	"context"
	"dia-manager-backend/config"
	"dia-manager-backend/utils"
	"errors"
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

func CreateToken(id string) (string, error) {
    return "", nil
}