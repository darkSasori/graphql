package repository

import (
	"context"

	"github.com/darksasori/graphql/pkg/model"
	"github.com/darksasori/graphql/pkg/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tweet repository
type Tweet struct {
	coll *mongo.Collection
}

// NewTweet return a pointer to tweet
func NewTweet(client *mongo.Database) *Tweet {
	return &Tweet{client.Collection("tweet")}
}

// FindByUser return all tweet user
func (t *Tweet) FindByUser(ctx context.Context, user *model.User) (service.Cursor, error) {
	return t.coll.Find(ctx, bson.D{{"user", user.Username}})
}

// FindAll return all tweet
func (t *Tweet) FindAll(ctx context.Context) (service.Cursor, error) {
	return t.coll.Find(ctx, bson.D{{}})
}

// FindOne return a tweet or error if not found
func (t *Tweet) FindOne(ctx context.Context, id interface{}) (*model.Tweet, error) {
	var result model.Tweet
	filter := bson.D{{"_id", id}}
	if err := t.coll.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Insert a new tweet in db
func (t *Tweet) Insert(ctx context.Context, tweet *model.Tweet) error {
	result, err := t.coll.InsertOne(ctx, tweet)
	if err != nil {
		return err
	}
	tweet.ID = result.InsertedID
	return nil
}

// Delete user
func (t *Tweet) Delete(ctx context.Context, tweet *model.Tweet) error {
	q := bson.M{"_id": tweet.ID}
	_, err := t.coll.DeleteOne(ctx, q)
	return err
}
