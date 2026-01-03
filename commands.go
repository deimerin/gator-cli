package main

import (
	"errors"
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
