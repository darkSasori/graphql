package service

import (
	"context"
	"fmt"

	"github.com/darksasori/graphql/model"
)

type User struct {
	userRepository UserRepository
}

func NewUser(repository UserRepository) *User {
	return &User{repository}
}

func (u *User) PrintList(ctx context.Context) error {
	var result interface{}
	cursor, err := u.userRepository.FindAll(ctx)
	if err != nil {
		return err
	}

	for cursor.Next(ctx) {
		err = cursor.Decode(&result)
		if err != nil {
			return err
		}
		fmt.Println(result)
	}

	return nil
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

func (u *User) Remove(ctx context.Context, user *model.User) error {
	return u.userRepository.Delete(ctx, user)
}
