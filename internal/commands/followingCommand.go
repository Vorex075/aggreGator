package commands

import (
	"context"
	"fmt"

	"github.com/Vorex075/aggreGator/internal/database"
)

func handleFollowing(s *State, _ Command, userInfo database.User) error {
	followingFeeds, err := s.db.GetFeedsFollowForUser(
		context.Background(), userInfo.Name)
	if err != nil {
		return err // The username does not exist. Impossible to happen.
	}
	fmt.Printf("'%s' is following:\n", s.cfg.CurrentUser)
	for _, feed := range followingFeeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}
	return nil
}
