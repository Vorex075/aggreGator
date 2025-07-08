package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Vorex075/aggreGator/internal/database"
	"github.com/Vorex075/aggreGator/internal/rss"
	"sync"
)

func scrapeFeeds(s *State) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	now := time.Now()

	fetchUpdate := database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  now,
			Valid: true,
		},
	}
	info, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		s.db.MarkFeedFetched(context.Background(), fetchUpdate)
		return err
	}
	var wg sync.WaitGroup
	ctx := context.Background()

	for _, item := range info.Channel.Item {

		wg.Add(1)
		go func(entry rss.RSSItem) {
			defer wg.Done()
			publicationTime, err := time.Parse(time.RFC1123Z, entry.PubDate)
			if err != nil {
				fmt.Printf("bad formatted publication time: %v\n", err)
				return
			}
			newPost := database.CreatePostParams{
				CreatedAt: now,
				UpdatedAt: now,
				Title:     entry.Title,
				Url:       entry.Link,
				Description: sql.NullString{
					Valid:  true,
					String: entry.Description,
				},
				PublishedAt: publicationTime,
				FeedID:      feed.ID,
			}
			_, err = s.db.CreatePost(ctx, newPost)
			if err != nil {
				fmt.Printf("error while creating a new post: %v\n", err)
			}
		}(item)
	}
	fmt.Printf("Fetching %d posts...\n", len(info.Channel.Item))
	wg.Wait()
	err = s.db.MarkFeedFetched(context.Background(), fetchUpdate)
	return err
}
