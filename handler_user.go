package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yhuet/aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("missing argument username")
	}
	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	err = s.conf.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("user set to: %s\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("missing argument username")
	}
	currentTime := time.Now()
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.args[0],
	})
	if err != nil {
		return err
	}
	err = s.conf.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("user created: %s\n", user.Name)
	fmt.Println(user)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Users:")
	for _, user := range users {
		if s.conf.CurrentUserName == user.Name {
			fmt.Println("*", user.Name, "(current)")

		} else {
			fmt.Println("*", user.Name)
		}
	}
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("users have been reset.")
	return nil
}
