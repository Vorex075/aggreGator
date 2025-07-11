package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Vorex075/aggreGator/internal/database"
)

func handleBookmark(s *State, cmd Command, userInfo database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("no post id provided in bookmar command")
	}
	value, err := strconv.Atoi(cmd.args[0])
	if err != nil {
		return err
	}
	queryParam := database.AddBookmarkParams{
		UserID: userInfo.ID,
		PostID: int32(value),
	}
	_, err = s.db.AddBookmark(context.Background(), queryParam)
	if err != nil {
		return err
	}

	fmt.Printf("Bookmark created on post with id - %d\n", value)
	return nil
}
