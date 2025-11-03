package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/yhuet/aggregator/internal/config"
	"github.com/yhuet/aggregator/internal/database"
)

type state struct {
	conf *config.Config
	db   *database.Queries
}

func main() {
	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening sql: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	state := state{
		conf: &cfg,
		db:   dbQueries,
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
