package main

import (
	"context"
	"errors"

	"github.com/deimerin/gator-cli/internal/database"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	fn, ok := c.cmds[cmd.name]
	if !ok {
		return errors.New("Command not found")
	}
	return fn(s, cmd)

}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return errors.New("user not logged in")
		}
		return handler(s, cmd, user)
	}
}
