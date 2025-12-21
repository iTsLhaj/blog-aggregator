package main

import (
	"context"

	"github.com/iTsLhaj/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) commandHandler {
	return func(s *state, cmd command) error {
		user, err := s.q.GetUser(context.Background(), s.c.Username)

		for _, cmdN := range []string{"reset", "register", "login"} {
			if cmd.name == cmdN {
				return handler(s, cmd, user)
			}
		}

		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
