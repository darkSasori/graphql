package utils

import (
	"context"

	"github.com/darksasori/graphql/pkg/model"
)

type key int

const userKey key = 0

func NewUserContext(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetUser(ctx context.Context) (*model.User, bool) {
	user, ok := ctx.Value(userKey).(*model.User)
	return user, ok
}
