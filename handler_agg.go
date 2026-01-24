package main

import (
	"context"
	"fmt"
	"time"
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

/*
Get the next feed to fetch from the DB.
Mark it as fetched.
Fetch the feed using the URL (we already wrote this function)
Iterate over the items in the feed and print their titles to the console.
*/

func scrapeFeeds(s *state) error {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		return fmt.Errorf("Couldn't get feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("Couldn't mark feed: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("Couldn't fetch feed: %w", err)
	}

	// Showing feed information
	fmt.Println("-------------------------------------------------")
	fmt.Printf("Feed Title: %s\n", rssFeed.Channel.Title)
	fmt.Printf("Feed description: %s\n", rssFeed.Channel.Description)
	fmt.Println("-------------------------------------------------")
	fmt.Println("Feed Items: ")

	for _, feed := range rssFeed.Channel.Item {
		fmt.Printf(" - %s\n", feed.Title)
	}
	fmt.Println("-------------------------------------------------")

	return nil
}
