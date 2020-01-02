package schema

import (
	"context"

	"github.com/darksasori/graphql/pkg/model"
	"github.com/darksasori/graphql/pkg/utils"
	"github.com/pkg/errors"
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

type userInput struct {
	Username, Displayname, Image *string
}

func (r *Resolver) SaveUser(ctx context.Context, args struct{ User userInput }) (*UserResolver, error) {
	u := &model.User{
		Username: *args.User.Username,
	}
	if args.User.Displayname != nil {
		u.Displayname = *args.User.Displayname
	}
	if args.User.Image != nil {
		u.Image = *args.User.Image
	}
	userLogger, ok := utils.GetUser(ctx)
	if !ok {
		return nil, errors.New("User not found")
	}
	if userLogger.Username != u.Username {
		return nil, errors.New("Operation is not permitted")
	}
	if err := r.user.Save(ctx, u); err != nil {
		return nil, errors.Wrap(err, "SaveUser")
	}

	return &UserResolver{u}, nil
}
