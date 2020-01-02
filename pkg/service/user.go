package service

import (
	"context"

	"github.com/darksasori/graphql/pkg/model"
)

// User Service
type User struct {
	Repository UserRepository
}

// NewUser return a pointer to User
func NewUser(repository UserRepository) *User {
	return &User{repository}
}

// Save the user
func (u *User) Save(ctx context.Context, user *model.User) error {
	userExist, err := u.Repository.FindOne(ctx, user.Username)
	if err != nil {
		if err := u.Repository.Insert(ctx, user); err != nil {
			return err
		}
		return nil
	}

	if user.Displayname != "" {
		userExist.Displayname = user.Displayname
	}

	if user.Image != "" {
		userExist.Image = user.Image
	}

	if err := u.Repository.Update(ctx, userExist); err != nil {
		return err
	}
	return nil
}

// Remove user
func (u *User) Remove(ctx context.Context, user *model.User) error {
	return u.Repository.Delete(ctx, user)
}
