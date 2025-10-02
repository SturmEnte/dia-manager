package auth

import (
	"context"
	"errors"
	"time"

	"dia-manager-backend/config"
	"dia-manager-backend/utils"

	"github.com/dchest/uniuri"
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

    token := uniuri.NewLen(40)
    expiresAt := time.Now().Add(5 * time.Minute).UTC()

    var dummy int

    err := config.DB.QueryRow(context.Background(), `INSERT INTO sessions (token, user_id, expires) VALUES ($1, $2, $3)`, token, id, expiresAt).Scan(&dummy)

    if err != nil && err.Error() != "no rows in result set" {
        println(err.Error())
        return "", errors.New("failed to insert session token")
    }

    return token, nil
}