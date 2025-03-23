package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Statement struct {
	Type    string
	Data    []string
	Success bool
}

func NewStatement() *Statement {
	return &Statement{
		Type:    "",
		Data:    make([]string, 0),
		Success: false,
	}
}

func main() {

	// Create and check if the database file exists
	_, err := os.Stat("database.db")
	var databaseFile *os.File
	if os.IsNotExist(err) {
		databaseFile, err = os.Create("database.db")
		if err != nil {
			fmt.Println("Error creating database storage: ", err)
		} else {
			fmt.Println("Database file storage created successfully")
		}
	} else {
		databaseFile, err = os.OpenFile("database.db", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Error loading database storage: ", err)
		} else {
			fmt.Println("Database file storage loaded successfully")
		}
	}

	reader := bufio.NewReader(os.Stdin)
	table := NewTable()
	databaseFile.Write(table.SerializeAllRows())

	// REPL
	for {

		statement := NewStatement()
		fmt.Print("gokulite> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) == 0 {
			fmt.Println("Error: Empty Statement")
			continue
		}

		// Handle the Meta command
		if input[0] == '.' {
			if input == ".exit" {
				fmt.Println("Exiting Gokulite. Goodbye!")
				break
			} else {
				err = executeMetaCommand(input)
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			// Handle the compiler command
			err := Compile(input, statement)
			if err != nil {
				fmt.Println("Compilation Error: ", err)
				continue
			}
			executeStatement(statement, table)
		}
	}

	defer databaseFile.Close()
}
