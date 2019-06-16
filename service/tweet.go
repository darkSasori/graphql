package service

import (
	"context"

	"github.com/darksasori/graphql/model"
)

// Tweet service
type Tweet struct {
	repository TweetRepository
}

// NewTweet return a pointer to Tweet
func NewTweet(repository TweetRepository) *Tweet {
	return &Tweet{repository}
}

// Save the tweet
func (t *Tweet) Save(ctx context.Context, tweet *model.Tweet) error {
	if err := t.repository.Insert(ctx, tweet); err != nil {
		return err
	}
	return nil
}

// Delete tweet
func (t *Tweet) Remove(ctx context.Context, tweet *model.Tweet) error {
	return t.repository.Delete(ctx, tweet)
}
