package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yhuet/aggregator/internal/database"
)

func handlerAddfeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("usage: addfeed <name> <url>")
	}
	currentTime := time.Now()
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("feed added:\n%+v\n", feed)

	res, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("User %s is following feed %s\n", res.UserName, res.FeedName)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	rows, err := s.db.GetFeedsWithUser(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Feeds:")
	for _, row := range rows {
		fmt.Printf(" - %s (%s) added by %s\n", row.FeedName, row.FeedUrl, row.UserName)
	}
	return nil
}
