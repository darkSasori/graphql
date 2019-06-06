package service

import (
	"context"

	"github.com/darksasori/graphql/model"
)

// User Service
type User struct {
	userRepository UserRepository
}

// NewUser return a pointer to User
func NewUser(repository UserRepository) *User {
	return &User{repository}
}

func (u *User) Save(ctx context.Context, user *model.User) error {
	if user.ID == nil {
		if err := u.userRepository.Insert(ctx, user); err != nil {
			return err
		}
		return nil
	}

	if _, err := u.userRepository.FindOne(ctx, user.ID); err != nil {
		return err
	}

	if err := u.userRepository.Update(ctx, user); err != nil {
		panic(err)
	}
	return nil
}

// Remove user
func (u *User) Remove(ctx context.Context, user *model.User) error {
	return u.userRepository.Delete(ctx, user)
}
