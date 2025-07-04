package commands

import (
	"context"
	"fmt"
)

// handleReset This function deletes all entries from the `users` table in the database.
func handleReset(s *State, cmd Command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Database successfully reseted!")
	return nil
}
