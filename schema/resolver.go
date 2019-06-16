package schema

import (
	"context"

	"github.com/darksasori/graphql/model"
	"github.com/darksasori/graphql/service"
	graphql "github.com/graph-gophers/graphql-go"
)

type Resolver struct {
	user *service.User
}

func New(user *service.User) *graphql.Schema {
	r := &Resolver{
		user,
	}
	return graphql.MustParseSchema(getSchema(), r)
}

func (r *Resolver) ListUsers(ctx context.Context) (*[]*UserResolver, error) {
	var users []*UserResolver

	cursor, err := r.user.Repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var u model.User
	for cursor.Next(ctx) {
		if err = cursor.Decode(&u); err != nil {
			return nil, err
		}
		users = append(users, &UserResolver{&u})
	}

	return &users, nil
}
