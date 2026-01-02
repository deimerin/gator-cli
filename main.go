package main

import (
	"fmt"

	"github.com/deimerin/gator-cli/internal/config"
)

func main() {
	// Read config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set current user to my name
	err = cfg.SetUser("deimerin")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the config file again and print the contents of the config struct to the terminal
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfg)

}
