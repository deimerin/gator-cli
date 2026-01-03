package main

import (
	"errors"
	"fmt"

	"github.com/deimerin/gator-cli/internal/config"
)

type state struct {
	cfg *config.Config
}

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return errors.New("The login handler expects a username argument")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("The user has been set")
	return nil

}
