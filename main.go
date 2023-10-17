package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(2)
	}

	switch os.Args[1] {
	case "command1":
		fmt.Println("Executing command 1")
		// Add your code for command1 here
	case "command2":
		fmt.Println("Executing command 2")
		// Add your code for command2 here
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(2)
	}
}
