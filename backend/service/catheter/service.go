package catheter

import (
	"context"
	"dia-manager-backend/config"
	"errors"
	"time"
)

func CreateCatheter(userId string, startedAt time.Time, endedAt *time.Time) (string, error) {

	var id string
	var err error
		
	println(startedAt.String())
	if endedAt != nil {
		println(endedAt.String())
	} else {
		println("endedAt is nil")
	}

	err = config.DB.QueryRow(context.Background(), `INSERT INTO catheters (user_id, started_at, ended_at) VALUES ($1, $2, $3) RETURNING id`, userId, startedAt, endedAt).Scan(&id)
	
	if err != nil {
	    println(err.Error())
	    return "", errors.New("failed to insert the new user into the database")
	}

	return id, nil
}