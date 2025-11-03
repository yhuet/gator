package main

import (
	"log"
	"os"

	"github.com/yhuet/aggregator/internal/config"
)

type state struct {
	conf *config.Config
}

func main() {
	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	state := state{
		conf: &cfg,
	}

	args := os.Args
	if len(args) < 2 {
		log.Fatal("not enough arguments")
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = commands.run(&state, cmd)
	if err != nil {
		log.Fatalf("error running command: %v", err)
	}
}
