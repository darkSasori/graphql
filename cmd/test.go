package main

import (
	"context"

	"github.com/darksasori/graphql/db"
	"github.com/darksasori/graphql/model"
	"github.com/darksasori/graphql/repository"
	"github.com/darksasori/graphql/service"
)

func main() {
	client, err := db.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUser(client.Database("mongodb-driver"))
	service := service.NewUser(userRepository)
	user := model.NewUser("lineufelipe", "Lineu Felipe")

	if err := service.Save(context.TODO(), user); err != nil {
		panic(err)
	}

	defer func() {
		if err := service.Remove(context.TODO(), user); err != nil {
			panic(err)
		}
	}()

	if err := service.PrintList(context.TODO()); err != nil {
		panic(err)
	}

	user.Displayname = "Lineuzinho"
	if err := service.Save(context.TODO(), user); err != nil {
		panic(err)
	}

	if err := service.PrintList(context.TODO()); err != nil {
		panic(err)
	}
}
