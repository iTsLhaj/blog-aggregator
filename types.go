package main

import (
	"database/sql"
	"errors"

	"github.com/iTsLhaj/gator/internal/config"
	"github.com/iTsLhaj/gator/internal/database"
)

type (
	state struct {
		db *sql.DB
		q  *database.Queries
		c  *config.Config
	}

	command struct {
		name string
		args []string
	}

	commands struct {
		cmdsList commandsRegister
	}

	commandHandler func(*state, command) error

	commandsRegister map[string]func(*state, command, database.User) error
)

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.cmdsList[cmd.name]
	if !ok {
		return errors.New("unknown command: " + cmd.name)
	}

	handler_ := middlewareLoggedIn(handler)
	err := handler_(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, handler func(*state, command, database.User) error) error {
	if _, ok := c.cmdsList[name]; ok {
		return errors.New("already registered command: " + name)
	}
	c.cmdsList[name] = handler
	return nil
}
