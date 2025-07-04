package commands

import (
	"context"
	"fmt"
)

func handleFeeds(s *State, cmd Command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("All feeds:")
	for i, feedEntry := range feeds {
		fmt.Printf("%d:\n", i)
		fmt.Printf("FeedName: %s\n", feedEntry.RssName)
		fmt.Printf("Url: %s\n", feedEntry.Url)
		fmt.Printf("UserName: %s\n", feedEntry.Username)
		fmt.Printf("\n\n-----------\n\n")
	}
	return nil
}
