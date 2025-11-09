package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
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
	log.Printf("Fetching posts from feed: %s\n", feed.Name)
	data, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	log.Printf("Found %d posts\n", len(data.Channel.Item))
	currentTime := time.Now()
	for _, item := range data.Channel.Item {
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Println(err)
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: publishedAt, Valid: true},
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("error: %+v\n", err)
		}
	}
	return nil
}
