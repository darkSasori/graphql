package schema

import (
	"context"

	"github.com/darksasori/graphql/pkg/model"
	"github.com/darksasori/graphql/pkg/service"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

type Resolver struct {
	user *service.User
}

func New(user *service.User) *graphql.Schema {
	r := &Resolver{
		user,
	}
	return graphql.MustParseSchema(schemaSctring, r)
}

func (r *Resolver) ListUsers(ctx context.Context) (*[]*UserResolver, error) {
	var users []*UserResolver

	cursor, err := r.user.Repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var u model.User
		if err = cursor.Decode(&u); err != nil {
			return nil, err
		}
		users = append(users, &UserResolver{&u})
	}

	return &users, nil
}

type userInput struct {
	Username, Displayname *string
}

func (r *Resolver) SaveUser(ctx context.Context, args struct{ User userInput }) (*UserResolver, error) {
	u := model.NewUser(*args.User.Username, *args.User.Displayname)
	if err := r.user.Save(ctx, u); err != nil {
		return nil, errors.Wrap(err, "SaveUser")
	}

	return &UserResolver{u}, nil
}
