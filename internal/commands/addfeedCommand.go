package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Vorex075/aggreGator/internal/database"
)

// handleAddfeed Adds a new feed into the database.
// Error cases are:
//
// - There are less than 2 arguments. Only the first two arguments will be considered.
//
// - The provided url already exists in the database
//
// - (nearly impossible) The returned value from the database could not be marshaled into a json string.
func handleAddfeed(s *State, cmd Command, userInfo database.User) error {
	argsLen := len(cmd.args)
	if argsLen < 2 {
		return fmt.Errorf("error on addfeed command: not enough arguments. Expected 2, found %d", argsLen)
	}
	now := time.Now()
	newFeed := database.CreateFeedParams{
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    userInfo.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	createdFeed, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return fmt.Errorf("error while creating the feed: %v", err)
	}
	fmt.Println("New feed created:")
	jsonData, _ := json.MarshalIndent(createdFeed, "", "\t")
	fmt.Println(string(jsonData))

	cmd.args[0] = newFeed.Url
	err = handleFollow(s, cmd, userInfo)
	return err
}
