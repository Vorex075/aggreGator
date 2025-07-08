package commands

import (
	"fmt"
	"os"
)

type Command struct {
	name string
	args []string
}

func CreateCommandFromArgs() (Command, error) {
	args := os.Args
	if len(args) < 2 {
		return Command{}, fmt.Errorf("no arguments provided")
	}
	return NewCommand(args[1], args[2:]), nil
}

func NewCommand(name string, args []string) Command {
	return Command{name: name, args: args}
}

type Commands struct {
	allowedCommands map[string]func(s *State, cmd Command) error
}

func NewCommands() Commands {
	cmds := Commands{allowedCommands: make(map[string]func(s *State, cmd Command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handleUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", middlewareLoggedIn(handleAddfeed))
	cmds.register("feeds", handleFeeds)
	cmds.register("follow", middlewareLoggedIn(handleFollow))
	cmds.register("following", middlewareLoggedIn(handleFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handleUnfollow))
	cmds.register("browse", middlewareLoggedIn(handleBrowse))
	return cmds
}

// run Tries to execute the command. If the command is not in the `commands` map, an error will be returned. If the command fail, the error is also returned.
func (c *Commands) Run(s *State, cmd Command) error {
	callback, ok := c.allowedCommands[cmd.name]
	if !ok {
		return fmt.Errorf("error while trying to execute a command: the command '%s' is invalid", cmd.name)
	}
	err := callback(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

// register Registers a new command into the `commands` struct. If the command already exists, it will be overwritten.
func (c *Commands) register(name string, f func(*State, Command) error) {
	c.allowedCommands[name] = f
}
