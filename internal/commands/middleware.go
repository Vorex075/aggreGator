package commands

import (
	"context"
	"fmt"

	"github.com/Vorex075/aggreGator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		userInfo, err := s.GetDB().GetUser(context.Background(), s.GetCurrentUser())
		if err != nil {
			return fmt.Errorf("User '%s' is not registered\n", s.GetCurrentUser())
		}
		return handler(s, cmd, userInfo)
	}
}
