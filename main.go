package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	compiler := NewCompiler()
	reader := bufio.NewReader(os.Stdin)

	// REPL 
	for { 
		fmt.Print("gokulite> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Handle the Meta command
		if input[0] == '.' {
			if input == ".exit" {
				fmt.Println("Exiting Gokulite. Goodbye!");
				break
			} else {
				fmt.Printf("Unrecognized Meta Command: %s\n", input);
			}
		} else {
			// Handle the compiler command
			program, err := compiler.Compile(input)
			if err != nil {
				fmt.Println("Compilation Error: ", err)
				continue
			}
			fmt.Printf("Unrecognized command: %s\n", input)
			fmt.Println("Compiled Program", program)
		}
	}
}