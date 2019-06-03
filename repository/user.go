package repository

import (
	"context"

	"github.com/darksasori/graphql/model"
	"github.com/darksasori/graphql/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	coll *mongo.Collection
}

func NewUser(client *mongo.Database) *User {
	return &User{client.Collection("user")}
}

func (u *User) getFilterOne(user *model.User) bson.M {
	return bson.M{"_id": user.ID}
}

func (u *User) Insert(ctx context.Context, user *model.User) error {
	result, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	user.ID = result.InsertedID
	return nil
}

func (u *User) FindAll(ctx context.Context) (service.Cursor, error) {
	return u.coll.Find(ctx, bson.D{{}})
}

func (u *User) FindOne(ctx context.Context, id interface{}) (*model.User, error) {
	var result model.User
	filter := bson.D{{"_id", id}}
	if err := u.coll.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (u *User) Delete(ctx context.Context, user *model.User) error {
	_, err := u.coll.DeleteOne(ctx, u.getFilterOne(user))
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(ctx context.Context, user *model.User) error {
	update := bson.D{
		{"$set", bson.D{
			{"username", user.Username},
			{"displayname", user.Displayname},
		}},
	}
	_, err := u.coll.UpdateOne(ctx, u.getFilterOne(user), update)
	if err != nil {
		return err
	}
	return nil
}
