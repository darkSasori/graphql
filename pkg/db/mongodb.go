package db

import (
	"context"
	"time"

	"github.com/darksasori/graphql/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {
	uri := utils.GetEnvDefault("MONGODB_URI", "")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	mongoClient = client
}

// Connect to db
func Connect(ctx context.Context) (*mongo.Client, error) {
	err := mongoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err = mongoClient.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return mongoClient, nil
}
