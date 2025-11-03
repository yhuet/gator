package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("missing argument username")
	}
	err := s.conf.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("user set to: %s\n", cmd.args[0])
	return nil
}
