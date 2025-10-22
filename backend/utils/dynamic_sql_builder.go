package utils

import (
	"strconv"
	"strings"

	"dia-manager-backend/types"
)

// Creates a dynamic insert sql statement based on the pairs which consist of
// variable-name and value
func BuildDynamicInsert(tableName string, varPairs []types.Pair, returning []string) (string, []interface{}) {

	var b strings.Builder
	b.WriteString("INSERT INTO ")
	b.WriteString(tableName)

	args := make([]interface{}, 0, len(varPairs))

	// Build var-name part
	b.WriteString(" (")

	for i := 0; i < len(varPairs); i++ {
		// This only executes if there is another pair and its not the first
		if(i > 0) {
			b.WriteString(",")
		}
		b.WriteString(varPairs[i].Key)
	}

	// Build value-part
	b.WriteString(") VALUES (")

	for i := 0; i < len(varPairs); i++ {
		// This only executes if there is another pair and its not the first
		if(i > 0) {
			b.WriteString(",")
		}
		b.WriteString("$")
		b.WriteString(strconv.Itoa(i+1))

		args = append(args, varPairs[i].Value)
	}

	// Build returning-part
	b.WriteString(") RETURNING ")

	for i := 0; i < len(returning); i++ {
		// This only executes if there is another pair and its not the first
		if(i > 0) {
			b.WriteString(",")
		}
		b.WriteString(returning[i])
	}

	return b.String(), args
}

// Creates a dynamic update sql statement based on the pairs which consist of
// variable-name and value
func BuildDynamicUpdate(tableName string, setPairs []types.Pair, wherePairs []types.Pair) (string, []interface{}) {

	var b strings.Builder
	b.WriteString("UPDATE ")
	b.WriteString(tableName)

	args := make([]interface{}, 0, len(setPairs)+len(wherePairs))

	// Build SET-part
	b.WriteString(" SET ")

	lastIndex := 0

	for i := 0; i < len(setPairs); i++ {
		// This only executes if there is another pair and its not the first
		if(i > 0) {
			b.WriteString(",")
		}
		b.WriteString(setPairs[i].Key)
		b.WriteString("=$")
		b.WriteString(strconv.Itoa(i+1))
		lastIndex = i+1

		args = append(args, setPairs[i].Value)
	}

	// Build SET-part
	b.WriteString(" WHERE ")

	for i := 0; i < len(wherePairs); i++ {
		// This only executes if there is another pair and its not the first
		if(i > 0) {
			b.WriteString(" AND ")
		}
		b.WriteString(wherePairs[i].Key)
		b.WriteString("=$")
		b.WriteString(strconv.Itoa(lastIndex + i+1))

		args = append(args, wherePairs[i].Value)
	}

	return b.String(), args
}
