package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Vorex075/aggreGator/internal/database"
	"github.com/Vorex075/aggreGator/internal/rss"
)

func scrapeFeeds(s *State) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	fetchUpdate := database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	info, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		s.db.MarkFeedFetched(context.Background(), fetchUpdate)
		return err
	}
	for _, item := range info.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}
	err = s.db.MarkFeedFetched(context.Background(), fetchUpdate)
	return err
}
