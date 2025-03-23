package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SaveData(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)

	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		return err
	}

	return fp.Sync()
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
	databaseFile.Write(table.Rows[0].Serialize().Bytes())

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
