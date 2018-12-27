package main

import (
	"fmt"
	"os"

	"github.com/baniol/awsp/profiles"
	"github.com/mitchellh/go-homedir"
)

func main() {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Error reading home directory!")
		os.Exit(1)
	}
	cred := fmt.Sprintf("%s/.aws/credentials", home)
	prof, err := profiles.NewProfiles(cred)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		prof.SetProfile()
		os.Exit(0)
	}

}
