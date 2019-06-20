package schema

import (
	"context"

	"github.com/darksasori/graphql/pkg/model"
)

type UserResolver struct {
	*model.User
}

func (u UserResolver) Displayname(ctx context.Context) *string {
	return &u.User.Displayname
}

func (u UserResolver) Username(ctx context.Context) *string {
	return &u.User.Username
}
