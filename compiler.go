package main

import (
	"fmt"
	"strings"
)

type Opcode int
const (
	OpInsert Opcode = iota
	OpSelect
)

type Instruction struct {
	Op   Opcode
	Args []interface{}
}

type Program struct {
	Instructions []Instruction
}

type Compiler struct{}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) Compile(input string, table *Table) (*Program, error) {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	program := &Program{}

	switch strings.ToLower(tokens[0]) {
	case "insert":
		if len(tokens) == 1 {
			return nil, fmt.Errorf("empty insert values")
		}

		table.Rows = append(table.Rows, Row{
			ID: tokens[1],
			Username: tokens[2],
			Email: tokens[3],
		})
		fmt.Println("Row updated successfully");
	
	case "select":
		fmt.Println("Select Command");

		if len(table.Rows) == 0 {
			fmt.Println("No rows found");
			break
		}
		for index, row := range table.Rows {
			fmt.Printf("[%d] ID: %s, Username: %s, Email: %s\n", 
            index+1,
            row.ID, 
            row.Username, 
            row.Email,
        )
		}

	case "delete":
		fmt.Println("Delete Command");
	default:
		return nil, fmt.Errorf("unrecognized command: %s", tokens[0]);
	}

	return program, nil
}

