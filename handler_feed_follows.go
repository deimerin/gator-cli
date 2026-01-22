package main

import (
	"context"
	"fmt"
	"time"

	"github.com/deimerin/gator-cli/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Missing or too many arguments.")
	}

	feedURL := cmd.args[0]
	currentUser := s.cfg.CurrentUserName

	feedByURL, err := s.db.GetFeedByURL(context.Background(), feedURL)

	if err != nil {
		return fmt.Errorf("Feed not found.")
	}

	userByName, err := s.db.GetUser(context.Background(), currentUser)

	if err != nil {
		return fmt.Errorf("Current username not found.")
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userByName.ID,
		FeedID:    feedByURL.ID,
	})

	if err != nil {
		return fmt.Errorf("Couldn't follow feed.")
	}

	fmt.Printf("%s now follows %s.\n", currentUser, feedByURL.Name)
	return nil

}

func handlerFollowing(s *state, cmd command) error {

	if len(cmd.args) != 0 {
		return fmt.Errorf("Too many arguments.")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Couldn't retrieve user. %w", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Couldn't retrieve feeds. %w", err)
	}

	// Priting user feeds
	fmt.Printf("%s following feeds:\n", user.Name)

	for _, feed := range follows {
		fmt.Printf(" - %s\n", feed.FeedName)
	}

	return nil
}
