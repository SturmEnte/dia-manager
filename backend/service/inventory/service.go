package inventory

import (
	"context"
	"encoding/json"
	"errors"

	"dia-manager-backend/env"
	"dia-manager-backend/types"
	"dia-manager-backend/utils"
)

func CreateItemStructure(userId string, name string, attributes map[string]interface{}) (string, error) {
	attributesJSON, err := json.Marshal(attributes)
	if err != nil {
		return "", errors.New("failed to marshal attributes")
	}

	pairs := []types.Pair{
		{Key: "user_id", Value: userId},
		{Key: "name", Value: name},
		{Key: "attributes", Value: attributesJSON},
	}

	query, args := utils.BuildDynamicInsert("item_structures", pairs, []string{"id"})

	var id string
	err = env.DB.QueryRow(context.Background(), query, args...).Scan(&id)
	if err != nil {
		return "", errors.New("failed to insert item structure into the database")
	}

	return id, nil
}