package service

import (
	"context"

	"github.com/darksasori/graphql/model"
)

// Cursor interface used in user service
type Cursor interface {
	Next(ctx context.Context) bool
	Decode(value interface{}) error
}

// UserRepository interface used to handle user in db
type UserRepository interface {
	FindAll(ctx context.Context) (Cursor, error)
	FindOne(ctx context.Context, id interface{}) (*model.User, error)
	Insert(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
}

// TweetRepository interface used to handle tweet in db
type TweetRepository interface {
	FindByUser(ctx context.Context, user *model.User) (Cursor, error)
	FindAll(ctx context.Context) (Cursor, error)
	FindOne(ctx context.Context, id interface{}) (*model.Tweet, error)
	Insert(ctx context.Context, tweet *model.Tweet) error
	Delete(ctx context.Context, tweet *model.Tweet) error
}
