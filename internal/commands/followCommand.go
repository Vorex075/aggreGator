package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/Vorex075/aggreGator/internal/database"
)

// handleFollow The active user (set in the *State param) will follow the feed specified in the first argument of the command. The function will return an error if:
//
// - There is no argument at all. Only 1 is requiered. Anyother is ignored.
// - The specified feed is not registered in the database.
func handleFollow(s *State, cmd Command, userInfo database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("error in follow command: not enough arguments. Expected 1, received 0")
	}
	url := cmd.args[0]

	/*userInfo, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return err // The user does not exists. Impossible to happen
	}*/
	feedInfo, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error in follow command: the '%s' feed is not registered yet\n", url)
	}

	now := time.Now()
	newFollowInfo := database.CreateFeedFollowParams{
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    userInfo.ID,
		FeedID:    feedInfo.ID,
	}
	insertedFollow, err := s.db.CreateFeedFollow(context.Background(), newFollowInfo)
	if err != nil {
		return fmt.Errorf("error in follow command: '%s' already follows '%s'", userInfo.Name, feedInfo.Name)
	}
	fmt.Printf("'%s' is now following '%s'\n", insertedFollow.Username, insertedFollow.FeedName)
	return nil
}
