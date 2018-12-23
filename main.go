package main

import (
	"fmt"
	"github.com/baniol/awsp/profiles"
	"os"
)

func main() {

	prof, err := profiles.NewProfiles()
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		prof.SetProfile()
		os.Exit(0)
	}

	command := os.Args[1]

	if command == "list" {
		prof.List()
	} else {
		fmt.Println("Command not found!")
		os.Exit(1)
	}

}
