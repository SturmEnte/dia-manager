package catheter

import (
	"context"
	"dia-manager-backend/config"
	"errors"
	"strconv"
	"time"
)

func CreateCatheter(userId string, startedAt time.Time, endedAt *time.Time) (string, error) {

	var id string

	err := config.DB.QueryRow(context.Background(), `INSERT INTO catheters (user_id, started_at, ended_at) VALUES ($1, $2, $3) RETURNING id`, userId, startedAt, endedAt).Scan(&id)
	
	if err != nil {
	    println(err.Error())
	    return "", errors.New("failed to insert the new user into the database")
	}

	return id, nil
}

func UpdateCatheter(userId string, catheterId string, startedAt *time.Time, endedAt *time.Time) (error) {

	query := "UPDATE catheters SET"
	args := []interface{}{}
	id := 1

	if startedAt != nil {

		query += " started_at=$" + strconv.Itoa(id) + ","
		args = append(args, startedAt)
		id++
	}

	if endedAt != nil {

		query += " ended_at=$" + strconv.Itoa(id)
		args = append(args, endedAt)
		id++
	}

	if len(args) == 0 {
		return errors.New("nothing to update")
	}

	// Missing not found error

	args = append(args, catheterId)
	args = append(args, userId)
	query += " WHERE id=$" + strconv.Itoa(id) + " AND user_id=$" + strconv.Itoa(id + 1) 

	var dummy int
	err := config.DB.QueryRow(context.Background(), query, args...).Scan(&dummy)
	
	if err != nil && err.Error() != "no rows in result set"  {
	    println(err.Error())
	    return errors.New("failed to insert the new user into the database")
	}

	return nil
}

func DeleteCatheter(userId string, catheterId string) (error) {

	var dummy int

	err := config.DB.QueryRow(context.Background(), `DELETE FROM catheters WHERE id=$1 AND user_id=$2`, catheterId, userId).Scan(&dummy)
	
	// Missing not found error

	if err != nil && err.Error() != "no rows in result set" {
	    println(err.Error())
	    return errors.New("failed to insert the new user into the database")
	}

	return nil
}