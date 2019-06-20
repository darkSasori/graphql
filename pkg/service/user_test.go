// +build test_unit

package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/darksasori/graphql/pkg/model"
	"gotest.tools/assert"
)

type cursor struct {
	users []*model.User
	pos   int
}

func (c *cursor) Next(ctx context.Context) bool {
	if c.pos >= len(c.users)-1 {
		return false
	}
	c.pos++
	return true
}

func (c *cursor) Decode(value interface{}) error {
	if c.pos >= len(c.users) {
		return fmt.Errorf("Invalid decode")
	}
	switch v := value.(type) {
	case *model.User:
		*v = *c.users[c.pos]
	default:
		return fmt.Errorf("Invalid interface")
	}
	return nil
}

type repository struct {
	users []*model.User
}

func (r *repository) FindAll(ctx context.Context) (Cursor, error) {
	return &cursor{r.users, -1}, nil
}

func (r *repository) FindOne(ctx context.Context, id interface{}) (*model.User, error) {
	for i := range r.users {
		if r.users[i].Username == id {
			return r.users[i], nil
		}
	}

	return nil, fmt.Errorf("User not found")
}

func (r *repository) Insert(ctx context.Context, user *model.User) error {
	r.users = append(r.users, user)
	return nil
}

func (r *repository) Update(ctx context.Context, user *model.User) error {
	for i := range r.users {
		if r.users[i].Username == user.Username {
			r.users[i].Displayname = user.Displayname
			return nil
		}
	}

	return fmt.Errorf("User not found")
}

func (r *repository) Delete(ctx context.Context, user *model.User) error {
	for i := range r.users {
		if r.users[i].Username == user.Username {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User not found")
}

func TestUserService(t *testing.T) {
	user := model.NewUser("username", "displayname")
	repo := &repository{}
	userService := NewUser(repo)

	if err := userService.Save(context.TODO(), user); err != nil {
		t.Error(err)
	}
	u := model.NewUser("username", "displayname")
	assert.Equal(t, u.Username, user.Username)

	user.Username = "username1"
	user.Displayname = "displayname1"
	if err := userService.Save(context.TODO(), user); err != nil {
		t.Error(err)
	}
	u = model.NewUser("username1", "displayname1")
	assert.Equal(t, u.Username, user.Username)
	assert.Equal(t, u.Displayname, user.Displayname)

	if err := userService.Remove(context.TODO(), user); err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(repo.users), 0)
}
