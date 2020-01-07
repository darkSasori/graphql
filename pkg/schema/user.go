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

func (u UserResolver) Image(ctx context.Context) *string {
	return &u.User.Image
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
