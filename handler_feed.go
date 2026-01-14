package main

import (
	"context"
	"fmt"
	"time"

	"github.com/deimerin/gator-cli/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Missing or too many arguments.\n")
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})

	if err != nil {
		return fmt.Errorf("can't create new feed: %v", err)
	}

	fmt.Printf("A new feed for %s has been created.\n", user.Name)

	fmt.Printf("ID: %v\nCreatedAt: %v\nUpdatedAt: %v\nName: %v\nURL: %v\nUserID: %v\n", feed.ID, feed.CreatedAt, feed.UpdatedAt, feed.Name, feed.Url, feed.UserID)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("can't retrieve feeds: %s", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed Name: %v\nFeed URL: %v\nFeed Creator: %v\n", feed.Feedname, feed.Feedurl, feed.Creatorusername)
	}

	return nil
}
