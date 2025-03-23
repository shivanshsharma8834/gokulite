package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Statement struct {
	statementType    string
	statementData    []any
	statementSuccess bool
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

	compiler := NewCompiler()
	reader := bufio.NewReader(os.Stdin)

	table := NewTable()
	table.Rows = append(table.Rows, Row{ID: "1", Username: "John", Email: "John@Gmail.com"})
	table.Rows = append(table.Rows, Row{ID: "2", Username: "John", Email: "John@Gmail.com"})
	table.Rows = append(table.Rows, Row{ID: "3", Username: "John", Email: "John@Gmail.com"})
	databaseFile.Write(table.SerializeAllRows())

	fileScanner := bufio.NewScanner(databaseFile)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	// REPL
	for {
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
			program, err := compiler.Compile(input)
			if err != nil {
				fmt.Println("Compilation Error: ", err)
				continue
			}
			fmt.Println("Compiled Program", program)
		}
	}

	defer databaseFile.Close()
}
