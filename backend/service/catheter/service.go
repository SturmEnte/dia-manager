package catheter

import (
	"context"
	"errors"
	"time"

	"dia-manager-backend/config"
	"dia-manager-backend/enums"
	"dia-manager-backend/models"
	"dia-manager-backend/types"
	"dia-manager-backend/utils"
)

func CreateCatheter(userId string, startedAt time.Time, endedAt *time.Time, changeReason *enums.ChangeReason) (string, error) {

	var id string

	pairs := []types.Pair{}

	pairs = append(pairs, types.Pair{Key: "user_id", Value: userId})
	pairs = append(pairs, types.Pair{Key: "started_at", Value: startedAt})

	if endedAt != nil {
		pairs = append(pairs, types.Pair{Key: "ended_at", Value: endedAt})
	}

	if changeReason != nil {
		pairs = append(pairs, types.Pair{Key: "change_reason", Value: changeReason})
	}

	query, args := utils.BuildDynamicInsert("catheters", pairs, []string{"id"})
	err := config.DB.QueryRow(context.Background(), query, args...).Scan(&id)

	if err != nil {
	    println(err.Error())
	    return "", errors.New("failed to insert the new user into the database")
	}

	return id, nil
}

func UpdateCatheter(userId string, catheterId string, startedAt *time.Time, endedAt *time.Time) (error) {
	setPairs := []types.Pair{}

	if startedAt != nil {
		setPairs = append(setPairs, types.Pair{Key: "started_at", Value: *startedAt})
	}

	if endedAt != nil {
		setPairs = append(setPairs, types.Pair{Key: "ended_at", Value: *endedAt})
	}

	if len(setPairs) == 0 {
		return errors.New("nothing to update")
	}

	wherePairs := []types.Pair{
		{Key: "id", Value: catheterId},
		{Key: "user_id", Value: userId},
	}

	query, args := utils.BuildDynamicUpdate("catheters", setPairs, wherePairs)

	println(query)

	var dummy int
	err := config.DB.QueryRow(context.Background(), query, args...).Scan(&dummy)
	
	if err != nil && err.Error() != "no rows in result set"  {
		println(err.Error())
		return errors.New("failed to update catheter in the database")
	}

	return nil
}

func GetCatheter(userId string, catheterId string) (models.Catheter, error) {
    
    var catheter models.Catheter

    err := config.DB.QueryRow(context.Background(), 
        `SELECT id, user_id, started_at, ended_at FROM catheters WHERE user_id=$1 AND id=$2`, 
        userId, catheterId).Scan(&catheter.ID, &catheter.UserID, &catheter.StartedAt, &catheter.EndedAt)

    if err != nil {
        if err.Error() == "no rows in result set" {
            // Return zero value for struct when not found
            return models.Catheter{}, errors.New("catheter not found")
        }
        println(err.Error())
        return models.Catheter{}, errors.New("failed to retrieve catheter from database")
    }

    return catheter, nil
}

func GetCatheters(userId string) ([]models.Catheter, error) {
	
	rows, err := config.DB.Query(context.Background(), `SELECT id, user_id, started_at, ended_at FROM catheters WHERE user_id = $1`, userId)
	
	if err != nil {
		println(err.Error())
		return nil, errors.New("failed to query catheters from the database")
	}
	
	defer rows.Close()

	// Create a slice to store the catheters
	var catheters []models.Catheter

	for rows.Next() {
		var catheter models.Catheter
		
		// Scan the row data into the catheter struct
		err := rows.Scan(&catheter.ID, &catheter.UserID, &catheter.StartedAt, &catheter.EndedAt)
		
		if err != nil {
			println(err.Error())
			return nil, errors.New("failed to scan catheter row")
		}
		
		// Append the catheter to the slice
		catheters = append(catheters, catheter)
	}

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		println(err.Error())
		return nil, errors.New("error occurred while iterating through catheter rows")
	}

	return catheters, nil
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