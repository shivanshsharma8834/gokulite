package main

import (
	"fmt"
	"strings"
)

func Compile(input string, statement *Statement) error {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return fmt.Errorf("empty input")
	}

	switch strings.ToLower(tokens[0]) {
	case "insert":
		if len(tokens) == 1 {
			return fmt.Errorf("empty insert values")
		}
		statement.Type = "insert"
		statement.Data = append(statement.Data, tokens[1:]...)

	case "select":
		statement.Type = "select"

	case "delete":
		statement.Type = "delete"
	default:
		return fmt.Errorf("unrecognized command: %s", tokens[0])
	}

	return nil
}
