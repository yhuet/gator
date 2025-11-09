package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yhuet/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	var err error
	if len(cmd.args) == 1 {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}
	if len(posts) == 0 {
		fmt.Println("No posts found")
	}
	for _, post := range posts {
		fmt.Printf("%+v\n", post)
	}
	return nil
}
