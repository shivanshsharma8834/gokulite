package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("gokulite> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == ".exit" {
			fmt.Println("Exiting Gokulite. Goodbye!")
			break
		}
		fmt.Printf("Unrecognized command: %s\n", input)
	}

	

}