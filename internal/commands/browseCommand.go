package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Vorex075/aggreGator/internal/database"
)

// handleBrowse Displays the most recent feed for the user.
// The only param is the limit.
//
// This command won't fail on execution, unless the database is down.
func handleBrowse(s *State, cmd Command, userInfo database.User) error {
	var limit int32 = 2
	if len(cmd.args) > 0 {
		var err error
		limitPreConverted, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(limitPreConverted)
	}
	queryParams := database.GetPostForUserParams{
		ID:    userInfo.ID,
		Limit: limit,
	}
	posts, err := s.db.GetPostForUser(context.Background(), queryParams)
	if err != nil {
		return err // Unexpected fail.
	}
	fmt.Printf("%d recent posts for %s:\n", len(posts), userInfo.Name)
	for _, post := range posts {
		fmt.Printf("* %s - Published at: %v\n", post.Url, post.PublishedAt)
	}
	return nil
}
