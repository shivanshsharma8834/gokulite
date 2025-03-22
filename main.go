package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
	"strings"
)

type Database struct {
	Name string
	Age  int
}

func writeDatatoFile(filename string, data any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	encoder := gob.NewEncoder(file)
	return encoder.Encode(data)
}

func readDatafromFile(filename string, data any) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	decoder := gob.NewDecoder(file)
	return decoder.Decode(&data)
}

func main() {

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

	writeDatatoFile("database.db", Database{Name: "John", Age: 32})
	var data any
	readDatafromFile("database.db", data)
	fmt.Println(data)

	compiler := NewCompiler()
	reader := bufio.NewReader(os.Stdin)
	table := NewTable()

	// REPL
	for {
		fmt.Print("gokulite> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Handle the Meta command
		if input[0] == '.' {
			if input == ".exit" {
				fmt.Println("Exiting Gokulite. Goodbye!")
				break
			} else {
				fmt.Printf("Unrecognized Meta Command: %s\n", input)
			}
		} else {
			// Handle the compiler command
			program, err := compiler.Compile(input, table)
			if err != nil {
				fmt.Println("Compilation Error: ", err)
				continue
			}
			fmt.Println("Compiled Program", program)
		}
	}

	defer databaseFile.Close()
}
