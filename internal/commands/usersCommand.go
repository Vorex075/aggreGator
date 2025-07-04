package commands

import (
	"context"
	"fmt"
)

// handleUsers Prints all users in the database, and marks out which one is
// currently logged in
func handleUsers(s *State, cmd Command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, usr := range users {
		fmt.Printf("* %s", usr.Name)
		if usr.Name == s.cfg.CurrentUser {
			fmt.Println(" (current)")
		} else {
			fmt.Println()
		}
	}
	return nil
}
