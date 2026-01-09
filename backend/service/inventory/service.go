package inventory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"dia-manager-backend/env"
	"dia-manager-backend/types"
	"dia-manager-backend/utils"
)

func GetItemStructures(userId string) ([]map[string]interface{}, error) {
	rows, err := env.DB.Query(context.Background(), "SELECT id, name, attributes FROM item_structures WHERE user_id=$1", userId)
	if err != nil {
		return nil, errors.New("failed to query item structures from database")
	}
	defer rows.Close()

	structures := make([]map[string]interface{}, 0)

	for rows.Next() {
		var id string
		var name string
		var attrsBytes []byte

		if err := rows.Scan(&id, &name, &attrsBytes); err != nil {
			return nil, errors.New("failed to scan item structure row")
		}

		var attributes map[string]map[string]interface{}
		if err := json.Unmarshal(attrsBytes, &attributes); err != nil {
			return nil, errors.New("failed to parse structure attributes from database")
		}

		structures = append(structures, map[string]interface{}{
			"id":         id,
			"name":       name,
			"attributes": attributes,
		})
	}

	if rows.Err() != nil {
		return nil, errors.New("error while iterating item structures rows")
	}

	return structures, nil
}

// ErrUserInput is returned when the client provided invalid data (bad JSON / invalid structure).
var ErrUserInput = errors.New("user input error")

func CreateItemStructure(userId string, name string, attributes map[string]interface{}) (string, error) {
	// Validate attributes: each parameter must be an object containing only
	// "required" (bool) and/or "default" (any). Anything else -> user error.
	for attrName, raw := range attributes {
		paramMap, ok := raw.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("%w: attribute '%s' must be an object with optional 'required' and/or 'default' fields", ErrUserInput, attrName)
		}

		for k := range paramMap {
			if k != "required" && k != "default" {
				return "", fmt.Errorf("%w: attribute '%s' contains invalid field '%s'", ErrUserInput, attrName, k)
			}
		}

		if req, exists := paramMap["required"]; exists {
			if _, ok := req.(bool); !ok {
				return "", fmt.Errorf("%w: attribute '%s' field 'required' must be boolean", ErrUserInput, attrName)
			}
		}
		// default can be any JSON value, no type check required
	}

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

func CreateItems(userId string, structureId string, items []map[string]interface{}) ([]string, error) {
	// Load structure attributes from DB
	var attrsBytes []byte
	err := env.DB.QueryRow(context.Background(), "SELECT attributes FROM item_structures WHERE id=$1", structureId).Scan(&attrsBytes)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: structure with id '%s' not found", ErrUserInput, structureId)
		}
		return nil, errors.New("failed to load item structure from database")
	}

	var structureAttrs map[string]map[string]interface{}
	if err := json.Unmarshal(attrsBytes, &structureAttrs); err != nil {
		return nil, errors.New("failed to parse structure attributes from database")
	}

	ids := make([]string, 0, len(items))

	for i := 0; i < len(items); i++ {
		item := items[i]

		// Verify that no item keys exist that are not defined in the structure
		for k := range item {
			if _, ok := structureAttrs[k]; !ok {
				return nil, fmt.Errorf("%w: item contains unknown parameter '%s'", ErrUserInput, k)
			}
		}

		// Ensure required parameters are present, and fill defaults for missing ones
		for paramName, paramDef := range structureAttrs {
			// paramDef may contain "required" (bool) and/or "default"
			required := false
			if r, ok := paramDef["required"]; ok {
				if rb, ok2 := r.(bool); ok2 {
					required = rb
				} else {
					return nil, fmt.Errorf("%w: structure definition for '%s' has invalid 'required' value", ErrUserInput, paramName)
				}
			}

			if _, provided := item[paramName]; !provided {
				// not provided in item
				if def, hasDefault := paramDef["default"]; hasDefault {
					// set default
					item[paramName] = def
				} else if required {
					return nil, fmt.Errorf("%w: required parameter '%s' missing", ErrUserInput, paramName)
				}
			}
		}

		// marshal item map to jsonb
		itemJSON, err := json.Marshal(item)
		if err != nil {
			return nil, errors.New("failed to marshal item data")
		}

		// print item json as string
		println(string(itemJSON))

		pairs := []types.Pair{
			{Key: "structure_id", Value: structureId},
			{Key: "data", Value: itemJSON},
		}

		query, args := utils.BuildDynamicInsert("items", pairs, []string{"id"})

		var id string
		err = env.DB.QueryRow(context.Background(), query, args...).Scan(&id)
		if err != nil {
			return nil, errors.New("failed to insert item into the database")
		}

		ids = append(ids, id)
	}

	return ids, nil
}