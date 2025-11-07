package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yhuet/aggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("usage: follow <url>")
	}

	user, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	currentTime := time.Now()
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

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Printf("User %s is not following any feed", user.Name)
		return nil
	}
	fmt.Printf("User %s is following:\n", user.Name)
	for _, feed := range feeds {
		fmt.Println(" -", feed.FeedName)
	}

	return nil
}
