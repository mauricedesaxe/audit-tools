package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(2)
	}

	switch os.Args[1] {
	case "logfiles":
		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		})
		if err != nil {
			fmt.Printf("Error walking the path %v: %v\n", ".", err)
			os.Exit(2)
		}
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(2)
	}
}
