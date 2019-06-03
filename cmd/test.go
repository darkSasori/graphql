package main

import (
	"context"
	"fmt"

	"github.com/darksasori/graphql/db"
	"github.com/darksasori/graphql/model"
	"github.com/darksasori/graphql/repository"
)

func main() {
	client, err := db.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUser(client.Database("mongodb-driver"))
	user := model.NewUser("lineufelipe", "Lineu Felipe")
	if err := userRepository.Insert(context.TODO(), &user); err != nil {
		panic(err)
	}
	fmt.Printf("Insert: %+v\n", user)

	cursor, err := userRepository.FindAll(context.TODO())
	if err != nil {
		panic(err)
	}

	var result model.User
	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&result); err != nil {
			panic(err)
		}
		fmt.Printf("FindAll: %+v\n", result)
	}

	user.Displayname = "Lineuzinho"
	if err = userRepository.Update(context.TODO(), &user); err != nil {
		panic(err)
	}
	fmt.Printf("Update: %+v\n", user)

	fmt.Println("Delete")
	if err = userRepository.Delete(context.TODO(), &user); err != nil {
		panic(err)
	}
}
