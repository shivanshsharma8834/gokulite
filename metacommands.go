package main

import "fmt"

func executeMetaCommand(command string) error {

	switch command {
	case ".help":
		fmt.Println("Help Desk")
		return nil

	default:
		return fmt.Errorf("unrecognized Meta Command: %s", command)
	}

}
