package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal xml
	var feedData RSSFeed
	err = xml.Unmarshal(data, &feedData)

	if err != nil {
		return nil, err
	}

	// UnscapeString
	feedData.Channel.Title = html.UnescapeString(feedData.Channel.Title)
	feedData.Channel.Description = html.UnescapeString(feedData.Channel.Description)

	for ind, feed := range feedData.Channel.Item {
		feedData.Channel.Item[ind].Title = html.UnescapeString(feed.Title)
		feedData.Channel.Item[ind].Description = html.UnescapeString(feed.Description)
	}

	return &feedData, nil

}
