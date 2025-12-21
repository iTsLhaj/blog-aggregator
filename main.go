package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var err error
	var st *state
	var cmds *commands

	if len(os.Args) < 2 {
		fmt.Println(errors.New("not enough arguments were provided"))
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	st, err = initState()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmdsRegister := commandsRegister{
		"login":     handlerLogin,
		"register":  handlerRegister,
		"reset":     handlerReset,
		"users":     handlerUsers,
		"agg":       handlerAgg,
		"addfeed":   handlerAddFeed,
		"feeds":     handlerFeeds,
		"follow":    handlerFollow,
		"following": handlerFollowing,
		"unfollow":  handlerUnfollow,
	}
	cmds, err = initCommands(cmdsRegister)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cmds.run(st, command{
		name: cmd,
		args: args,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
