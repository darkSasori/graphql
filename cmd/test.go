package main

import (
	"context"

	"github.com/darksasori/graphql/pkg/db"
	"github.com/darksasori/graphql/pkg/model"
	"github.com/darksasori/graphql/pkg/repository"
	"github.com/darksasori/graphql/pkg/service"
	"github.com/darksasori/graphql/pkg/utils"
)

func main() {
	conn, err := db.Connect(context.TODO())
	if err != nil {
		panic(err)
	}
	connDB := conn.Database(utils.GetEnvDefault("MONGODB_NAME", "blog"))

	userService := service.NewUser(repository.NewUser(connDB))
	user := model.NewUser("lineufelipe", "Lineu Felipe")
	if err := userService.Save(context.TODO(), user); err != nil {
		panic(err)
	}

	tweetService := service.NewTweet(repository.NewTweet(connDB))
	tweet := model.NewTweet("just a test", user)
	if err := tweetService.Save(context.TODO(), tweet); err != nil {
		panic(err)
	}
}
