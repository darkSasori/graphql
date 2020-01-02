package repository

import (
	"context"

	"github.com/darksasori/graphql/pkg/model"
	"github.com/darksasori/graphql/pkg/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User repository
type User struct {
	coll *mongo.Collection
}

// NewUser return a pointer to user
func NewUser(client *mongo.Database) *User {
	return &User{client.Collection("user")}
}

func (u *User) getFilterOne(user *model.User) bson.M {
	return bson.M{"_id": user.Username}
}

// Insert a new user in db
func (u *User) Insert(ctx context.Context, user *model.User) error {
	if _, err := u.coll.InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}

// FindAll return all user
func (u *User) FindAll(ctx context.Context) (service.Cursor, error) {
	return u.coll.Find(ctx, bson.D{{}})
}

// FindOne return a user or error if not found
func (u *User) FindOne(ctx context.Context, id interface{}) (*model.User, error) {
	var result model.User
	filter := bson.D{{"_id", id}}
	if err := u.coll.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete user
func (u *User) Delete(ctx context.Context, user *model.User) error {
	_, err := u.coll.DeleteOne(ctx, u.getFilterOne(user))
	if err != nil {
		return err
	}
	return nil
}

// Update user
func (u *User) Update(ctx context.Context, user *model.User) error {
	update := bson.D{
		{"$set", bson.D{
			{"displayname", user.Displayname},
			{"image", user.Image},
		}},
	}
	_, err := u.coll.UpdateOne(ctx, u.getFilterOne(user), update)
	if err != nil {
		return err
	}

	return nil
}
