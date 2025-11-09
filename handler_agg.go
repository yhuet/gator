package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/yhuet/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("usage: agg <duration>")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        feed.ID,
	})
	if err != nil {
		return err
	}
	data, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	fmt.Printf("feed %s:\n", data.Channel.Title)
	for i, item := range data.Channel.Item {
		fmt.Println(item.Title, "<", item.Link, ">")
		if i > 8 {
			break
		}
	}
	fmt.Println("==============================")
	return nil
}
