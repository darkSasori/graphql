package main

import (
	"context"

	"github.com/darksasori/graphql/db"
	"github.com/darksasori/graphql/model"
	"github.com/darksasori/graphql/repository"
	"github.com/darksasori/graphql/service"
)

func main() {
	conn, err := db.Connect(context.TODO())
	if err != nil {
		panic(err)
	}
	connDB := conn.Database("test")

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
