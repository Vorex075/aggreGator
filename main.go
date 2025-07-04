package main

import (
	"fmt"
	"log"

	"github.com/Vorex075/aggreGator/internal/commands"
	"github.com/Vorex075/aggreGator/internal/config"
	_ "github.com/lib/pq"
) // We import it for its side effects.
// This is kind of mandatory to work with sql in golang...

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	state := commands.NewState(&cfg)
	cmds := commands.NewCommands()
	command, err := commands.CreateCommandFromArgs()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = cmds.Run(state, command)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}
