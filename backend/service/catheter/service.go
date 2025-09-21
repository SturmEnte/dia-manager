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

func UpdateCatheter(catheterId string, startedAt *time.Time, endedAt *time.Time) (error) {

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

	args = append(args, catheterId)
	query += " WHERE id=$" + strconv.Itoa(id)

	println(query)

	var dummy int
	err := config.DB.QueryRow(context.Background(), query, args...).Scan(&dummy)
	
	if err != nil && err.Error() != "no rows in result set"  {
	    println(err.Error())
	    return errors.New("failed to insert the new user into the database")
	}

	return nil
}

func DeleteCatheter(catheterId string) (error) {

	var dummy int

	err := config.DB.QueryRow(context.Background(), `DELETE FROM catheters WHERE id=$1`, catheterId).Scan(&dummy)
	
	if err != nil && err.Error() != "no rows in result set" {
	    println(err.Error())
	    return errors.New("failed to insert the new user into the database")
	}

	return nil
}