package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/iTsLhaj/gator/internal/config"
	"github.com/iTsLhaj/gator/internal/database"

	_ "github.com/lib/pq"
)

func initState() (*state, error) {
	var state state

	cfg, _ := config.Read()
	state.c = &cfg

	db, err := sql.Open("postgres", state.c.DbUrl)
	if err != nil {
		return nil, err
	}
	state.db = db
	state.q = database.New(db)

	return &state, nil
}

func initCommands(cmdReg commandsRegister) (*commands, error) {
	var cmds commands = commands{
		cmdsList: make(commandsRegister),
	}

	for name, handler := range cmdReg {
		err := cmds.register(name, handler)
		if err != nil {
			return nil, err
		}
	}

	return &cmds, nil
}

func feedListPrettier(s *state, fl []database.Feed) error {
	biggestNameLenght := 0
	biggestURLLenght := 0
	biggestUserNameLenght := 0

	repeat := func(s string, n int) string {
		return strings.Repeat(s, n)
	}

	pad := func(s string, width int) string {
		return s + repeat(" ", width-len(s))
	}

	max_ := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, feed := range fl {
		if biggestNameLenght < len(feed.Name) {
			biggestNameLenght = len(feed.Name)
		}
		if biggestURLLenght < len(feed.Url) {
			biggestURLLenght = len(feed.Url)
		}
		user, err := s.q.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		if biggestUserNameLenght < len(user.Name) {
			biggestUserNameLenght = len(user.Name)
		}
	}

	padding := 3

	nameW := max_(biggestNameLenght, len("Name")) + padding*2
	urlW := max_(biggestURLLenght, len("URL")) + padding*2
	userW := max_(biggestUserNameLenght, len("User")) + padding*2

	// top border
	fmt.Printf("┌%s┬%s┬%s┐\n",
		repeat("─", nameW),
		repeat("─", urlW),
		repeat("─", userW),
	)

	// header
	fmt.Printf("│ %s  │ %s  │ %s  │\n",
		pad("Name", nameW-padding),
		pad("URL", urlW-padding),
		pad("User", userW-padding),
	)

	// separator
	fmt.Printf("├%s┼%s┼%s┤\n",
		repeat("─", nameW),
		repeat("─", urlW),
		repeat("─", userW),
	)

	// rows
	for _, feed := range fl {
		user, _ := s.q.GetUserByID(context.Background(), feed.UserID)

		fmt.Printf("│ %s  │ %s  │ %s  │\n",
			pad(feed.Name, nameW-padding),
			pad(feed.Url, urlW-padding),
			pad(user.Name, userW-padding),
		)
	}

	// bottom border
	fmt.Printf("└%s┴%s┴%s┘\n",
		repeat("─", nameW),
		repeat("─", urlW),
		repeat("─", userW),
	)

	return nil
}
