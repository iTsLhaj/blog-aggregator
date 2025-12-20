package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iTsLhaj/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("login <username>")
	}

	user, err := s.q.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.c.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("login success")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("register <username>")
	}

	user, err := s.q.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	err = s.c.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("register success")
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("`reset` has no arguments")
	}

	err := s.q.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("reset success")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("`users` has no arguments")
	}

	users, err := s.q.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Printf("* %s ", user.Name)
		if s.c.Username == user.Name {
			fmt.Printf("(current)")
		}
		fmt.Printf("\n")
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("`agg` has no arguments")
	}

	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("addfeed <title> <url>")
	}

	var userName string = s.c.Username
	var feedTitle string = cmd.args[0]
	var feedUrl string = cmd.args[1]

	_, err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	var user database.User
	user, err = s.q.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}

	var feed database.Feed
	feed, err = s.q.AddFeed(
		context.Background(),
		database.AddFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      feedTitle,
			Url:       feedUrl,
			UserID:    user.ID,
		},
	)
	if err != nil {
		return err
	}

	_, err = s.q.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("`feeds` has no arguments")
	}

	feeds, err := s.q.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	err = feedListPrettier(s, feeds)
	if err != nil {
		return err
	}

	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("follow <url>")
	}

	var err error
	var user database.User
	var feed database.Feed

	user, err = s.q.GetUser(context.Background(), s.c.Username)
	if err != nil {
		return err
	}

	feed, err = s.q.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	var cffr database.CreateFeedFollowRow
	cffr, err = s.q.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("follow success")
	fmt.Printf(" - user: %s\n", cffr.UserName)
	fmt.Printf(" - feed: %s\n", cffr.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("`following` has no arguments")
	}

	user, err := s.q.GetUser(context.Background(), s.c.Username)
	if err != nil {
		return err
	}

	var feedFollows []database.GetFeedFollowsForUserRow
	feedFollows, err = s.q.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, ff := range feedFollows {
		fmt.Printf("%s\n", ff.FeedName)
	}

	return nil
}
