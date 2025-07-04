package commands

import (
	"context"
	"fmt"

	"github.com/Vorex075/aggreGator/internal/database"
)

func handleUnfollow(s *State, cmd Command, userInfo database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("error in unfollow command: no url specified")
	}
	feedInfo, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("error in unfollow command: the specified feed url is not registered")
	}
	err = s.db.UnfollowFeed(context.Background(),
		database.UnfollowFeedParams{
			UserID: userInfo.ID,
			FeedID: feedInfo.ID,
		})
	if err != nil {
		return err // This should never be reached.
	}
	fmt.Printf("'%s' doesn't follow '%s' anymore", userInfo.Name, feedInfo.Name)
	return nil
}
