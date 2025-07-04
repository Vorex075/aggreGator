package commands

import (
	"context"
	"fmt"
)

// handlerLogin Modifies the active user. If the user does not exist int the database,
// an error will be returned.
func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("error when executing 'login' command: no arguments provided")
	}
	usrName := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), usrName)
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(usrName)
	if err != nil {
		return err
	}
	fmt.Printf("User set to: %s\n", usrName)
	return nil
}
