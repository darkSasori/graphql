package schema

import (
	"context"

	"github.com/darksasori/graphql/pkg/model"
	graphql "github.com/graph-gophers/graphql-go"
)

type UserResolver struct {
	*model.User
}

func (UserResolver) ID(ctx context.Context) *graphql.ID {
	s := graphql.ID("")
	return &s
}

func (u UserResolver) Displayname(ctx context.Context) *string {
	return &u.User.Displayname
}

func (u UserResolver) Username(ctx context.Context) *string {
	return &u.User.Username
}
