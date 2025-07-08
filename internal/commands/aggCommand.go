package commands

import (
	"fmt"
	"time"
)

// handleAgg Prints out the information of a rss
func handleAgg(s *State, cmd Command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("error in agg command: you need to specify the time between fetch\n")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error in agg command: %v\n", err)
	}

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Printf("error in agg command: %v\n", err)
		}
	}
}
