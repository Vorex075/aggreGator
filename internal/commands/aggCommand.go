package commands

import (
	"context"
	"fmt"

	"github.com/Vorex075/aggreGator/internal/rss"
)

// handleAgg Prints out the information of a rss
func handleAgg(s *State, cmd Command) error {
	data, err := rss.FetchFeed(context.Background(), "https://wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(*data)

	return nil
}
