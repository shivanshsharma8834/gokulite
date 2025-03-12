package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Row struct {
	ID       int
	Username string
	Email    string
}

type Table struct {
	Rows []Row
}

func NewTable() *Table {
	return &Table{
		Rows: make([]Row, 0),
	}
}

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

func (c *Compiler) Compile(input string) (*Program, error) {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return nil, fmt.Errorf("Empty Input")
	}

	program := &Program{}

	switch strings.ToLower(tokens[0]) {
	case "insert":
		id, _ := strconv.Atoi(tokens[1])
		program.Instructions = append(program.Instructions, Instruction{
			Op: OpInsert,
			Args: []interface{}{id, tokens[2], tokens[3]},
		})
	case "select":
		program.Instructions = append(program.Instructions, Instruction{
			Op: OpSelect,
		})

	default:
		return nil, fmt.Errorf("Unrecognized command: %s", tokens[0])
	}

	return program, nil
}

