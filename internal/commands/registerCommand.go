package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/Vorex075/aggreGator/internal/database"
	"github.com/google/uuid"
)

// handlerRegister Registers a new user. Returns an error if no name is provided
// (args < 1), if the user name already exists in the database or if there was
// an error while modifying the config file (check config.SetUser() function).
func handlerRegister(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no name provided in register command")
	}
	now := time.Now()
	newName := cmd.args[0]
	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      newName,
	}
	_, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return err
	}
	fmt.Println("User successfully created!")
	err = s.cfg.SetUser(newName)
	if err != nil {
		return err
	}
	fmt.Printf("Current user: %s\n", newName)
	return nil
}
