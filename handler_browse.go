package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/deimerin/gator-cli/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	postLimit := 2
	if len(cmd.args) == 1 {
		limit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		postLimit = limit
	}

	ctx := context.Background()

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(postLimit),
	}

	posts, err := s.db.GetPostsForUser(ctx, params)
	if err != nil {
		return err
	}

	// Posts
	for _, post := range posts {

		fmt.Printf("[%s] %s\n", post.FeedName, post.Title)
		fmt.Printf("  URL: %s\n", post.Url)

		if post.PublishedAt.Valid {
			fmt.Printf("  Published: %s\n", post.PublishedAt.Time)
		} else {
			fmt.Printf("  Published: (unknown)\n")
		}
		fmt.Println()
	}
	return nil
}
