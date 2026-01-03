package main

import (
	"fmt"
	"os"

	"github.com/deimerin/gator-cli/internal/config"
)

func main() {
	// read config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Cant read the config file")
		os.Exit(1)
	}

	// set state
	st := state{
		cfg: &cfg,
	}

	// make command struct and register login function
	commandList := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	commandList.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments. args < 2")
		os.Exit(1)
	}

	// new command from arguments
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	// run command

	err = commandList.run(&st, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
