package commands

import (
	"context"
	"fmt"
	"time"

	"strconv"

	"github.com/Vorex075/aggreGator/internal/database"
)

// handleBrowse Displays the most recent feed for the user.
// The only param is the limit.
//
// This command won't fail on execution, unless the database is down.
func handleBrowse(s *State, cmd Command, userInfo database.User) error {
	var limit int32 = 2
	if len(cmd.args) >= 1 {
		readLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(readLimit)
	}
	queryParams2 := database.GetRecentPostsForUserParams{
		UserID:      userInfo.ID,
		PublishedAt: s.cfg.LastPost.Publicated_at,
		Limit:       limit,
	}
	posts, err := s.db.GetRecentPostsForUser(context.Background(), queryParams2)
	if err != nil {
		return err
	}
	if len(posts) == 0 {
		fmt.Println("No recent posts")
		s.cfg.UpdateLastPost(time.Now(), 0)
		return nil
	}
	for _, post := range posts {
		fmt.Printf("* %s - ID: %v\n", post.Title, post.ID)
	}
	s.cfg.UpdateLastPost(posts[len(posts)-1].PublishedAt, int(posts[len(posts)-1].ID))
	return nil
}
