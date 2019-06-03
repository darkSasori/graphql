package service

import (
	"context"

	"github.com/darksasori/graphql/model"
)

type Cursor interface {
	Next(ctx context.Context) bool
	Decode(value interface{}) error
}

type UserRepository interface {
	FindAll(ctx context.Context) (Cursor, error)
	FindOne(ctx context.Context, id interface{}) (*model.User, error)
	Insert(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
}
