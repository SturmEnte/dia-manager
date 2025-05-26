package utils

import (
	"context"
	"dia-manager-backend/config"
	"errors"
)

func GetUserIdByToken(token string) (string, error) {

	var userId string

	err := config.DB.QueryRow(context.Background(), `SELECT user_id FROM sessions WHERE token=$1`, token).Scan(&userId)

	if err != nil {
		println(err.Error())
		return "", errors.New("failed to insert session token")
	}

	return userId, nil
}