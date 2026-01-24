package inventory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"dia-manager-backend/env"
	"dia-manager-backend/types"
	"dia-manager-backend/utils"
)

func GetItemStructures(userId string) ([]map[string]interface{}, error) {
	rows, err := env.DB.Query(context.Background(), "SELECT id, name, attributes, general_information FROM item_structures WHERE user_id=$1", userId)
	if err != nil {
		return nil, errors.New("failed to query item structures from database")
	}
	defer rows.Close()

	structures := make([]map[string]interface{}, 0)

	for rows.Next() {
		var id string
		var name string
		var attrsBytes []byte
		var generalInfoBytes []byte

		if err := rows.Scan(&id, &name, &attrsBytes, &generalInfoBytes); err != nil {
			return nil, errors.New("failed to scan item structure row")
		}

		var attributes []map[string]interface{}
		if err := json.Unmarshal(attrsBytes, &attributes); err != nil {
			return nil, errors.New("failed to parse structure attributes from database")
		}

		var generalInformation []map[string]interface{}
		if err := json.Unmarshal(generalInfoBytes, &generalInformation); err != nil {
			return nil, errors.New("failed to parse structure general information from database")
		}

		structures = append(structures, map[string]interface{}{
			"id":                   id,
			"name":                 name,
			"attributes":           attributes,
			"general_information":  generalInformation,
		})
	}

	if rows.Err() != nil {
		return nil, errors.New("error while iterating item structures rows")
	}

	return structures, nil
}

var ErrUserInput = errors.New("user input error")

func CreateItemStructure(userId string, name string, attributes []map[string]interface{}, generalInformation []map[string]interface{}) (string, error) {
	// Validate attributes: each attribute must be an object containing a "required" field (bool)
	// and may contain a "name" field (string)
	for i, attr := range attributes {
		if attr == nil {
			return "", fmt.Errorf("%w: attribute at index %d is null", ErrUserInput, i)
		}

		// Check if "required" field exists and is boolean
		if req, exists := attr["required"]; exists {
			if _, ok := req.(bool); !ok {
				return "", fmt.Errorf("%w: attribute at index %d field 'required' must be boolean", ErrUserInput, i)
			}
		} else {
			return "", fmt.Errorf("%w: attribute at index %d must have a 'required' field", ErrUserInput, i)
		}

		// Validate allowed fields: name, required
		for k := range attr {
			if k != "name" && k != "required" {
				return "", fmt.Errorf("%w: attribute at index %d contains invalid field '%s'", ErrUserInput, i, k)
			}
		}
	}

	// Validate general information: each entry can contain any fields
	// No specific validation required for general information entries

	attributesJSON, err := json.Marshal(attributes)
	if err != nil {
		return "", errors.New("failed to marshal attributes")
	}

	generalInfoJSON, err := json.Marshal(generalInformation)
	if err != nil {
		return "", errors.New("failed to marshal general information")
	}

	pairs := []types.Pair{
		{Key: "user_id", Value: userId},
		{Key: "name", Value: name},
		{Key: "attributes", Value: attributesJSON},
		{Key: "general_information", Value: generalInfoJSON},
	}

	query, args := utils.BuildDynamicInsert("item_structures", pairs, []string{"id"})

	var id string
	err = env.DB.QueryRow(context.Background(), query, args...).Scan(&id)
	if err != nil {
		return "", errors.New("failed to insert item structure into the database")
	}

	return id, nil
}

func GetItems(userId string) ([]map[string]interface{}, error) {
	rows, err := env.DB.Query(context.Background(), "SELECT items.id, items.data, items.structure_id, items.created_at FROM items JOIN item_structures ON structure_id=item_structures.id WHERE item_structures.user_id=$1 ORDER BY items.created_at", userId)
	if err != nil {
		return nil, errors.New("failed to query items from database")
	}
	defer rows.Close()

	items := make([]map[string]interface{}, 0)

	for rows.Next() {
		var id string
		var dataBytes []byte
		var structureId string
		var createdAt time.Time

		if err := rows.Scan(&id, &dataBytes, &structureId, &createdAt); err != nil {
			return nil, errors.New("failed to scan item row")
		}

		var data map[string]interface{}
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			return nil, errors.New("failed to parse item data from database")
		}

		items = append(items, map[string]interface{}{
			"id":           id,
			"structure_id": structureId,
			"data":         data,
			"created_at":   createdAt,
		})
	}

	if rows.Err() != nil {
		return nil, errors.New("error while iterating item rows")
	}

	return items, nil
}

func CreateItems(userId string, structureId string, items []map[string]interface{}) ([]string, error) {
	// Load structure attributes from DB
	var attrsBytes []byte
	err := env.DB.QueryRow(context.Background(), "SELECT attributes FROM item_structures WHERE id=$1 AND user_id=$2", structureId, userId).Scan(&attrsBytes)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: structure with id '%s' not found", ErrUserInput, structureId)
		}
		return nil, errors.New("failed to load item structure from database")
	}

	var structureAttrs []map[string]interface{}
	if err := json.Unmarshal(attrsBytes, &structureAttrs); err != nil {
		return nil, errors.New("failed to parse structure attributes from database")
	}

	// Build a map of attribute names to their definitions for easier lookup
	attrMap := make(map[string]map[string]interface{})
	for _, attr := range structureAttrs {
		if name, ok := attr["name"].(string); ok {
			attrMap[name] = attr
		}
	}

	ids := make([]string, 0, len(items))

	for i := 0; i < len(items); i++ {
		item := items[i]

		// Verify that no item keys exist that are not defined in the structure
		for k := range item {
			if _, ok := attrMap[k]; !ok {
				return nil, fmt.Errorf("%w: item contains unknown parameter '%s'", ErrUserInput, k)
			}
		}

		// Ensure required parameters are present
		for attrName, paramDef := range attrMap {
			// paramDef may contain "required" (bool) and "name" (string)
			required := false
			if r, ok := paramDef["required"]; ok {
				if rb, ok2 := r.(bool); ok2 {
					required = rb
				} else {
					return nil, fmt.Errorf("%w: structure definition for '%s' has invalid 'required' value", ErrUserInput, attrName)
				}
			}

			if _, provided := item[attrName]; !provided {
				// not provided in item
				if required {
					return nil, fmt.Errorf("%w: required parameter '%s' missing", ErrUserInput, attrName)
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