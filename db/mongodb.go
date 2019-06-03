package db

import (
	"context"
	"time"

	"github.com/darksasori/graphql/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MONGODB_URI))
	if err != nil {
		panic(err)
	}

	mongoClient = client
}

func Connect(ctx context.Context) (*mongo.Client, error) {
	err := mongoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if err = mongoClient.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return mongoClient, nil
}
