package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/deimerin/gator-cli/internal/config"
	"github.com/deimerin/gator-cli/internal/database"
	"github.com/google/uuid"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return errors.New("The login handler expects a username argument")
	}

	userName := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return errors.New("could not find user.")
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Println("The user has been set")
	return nil

}

func handlerRegister(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return errors.New("expected name")
	}

	newUser := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), newUser)
	if err == nil {
		return errors.New("user already exists")
	}

	_, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      newUser,
	})

	if err != nil {
		return errors.New("something went wrong creating a new user")
	}

	err = s.cfg.SetUser(newUser)
	if err != nil {
		return errors.New("cannot set new user as current user")
	}

	fmt.Printf("user created: %s\n", newUser)
	return nil

}

func handleReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return errors.New("couldn't reset users table")
	}
	fmt.Println("users table has been reset")
	return nil
}
