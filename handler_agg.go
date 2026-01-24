package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/deimerin/gator-cli/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("Missing time interval argument.")
	}

	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Couldn't parse duration from argument: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Printf("Couldn't retrieve feed: %s\n", err)
		}
	}

}

func scrapeFeeds(s *state) error {

	ctx := context.Background()

	nextFeed, err := s.db.GetNextFeedToFetch(ctx)

	if err != nil {
		return fmt.Errorf("Couldn't get feed: %w", err)
	}

	err = s.db.MarkFeedFetched(ctx, nextFeed.ID)
	if err != nil {
		return fmt.Errorf("Couldn't mark feed: %w", err)
	}

	rssFeed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return fmt.Errorf("Couldn't fetch feed: %w", err)
	}

	for _, item := range rssFeed.Channel.Item {

		desc := sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}

		var pub sql.NullTime
		if item.PubDate != "" {
			t, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err == nil {
				pub = sql.NullTime{
					Time:  t,
					Valid: true,
				}
			}

		}

		_, err := s.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: desc,
			PublishedAt: pub,
			FeedID:      nextFeed.ID,
		})

		if err != nil {

			var pgErr *pq.Error
			if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" { // 23505 = unique_violation
					continue
				}
			}
			fmt.Println("error creating post entry.")
			continue
		}

	}

	return nil
}
