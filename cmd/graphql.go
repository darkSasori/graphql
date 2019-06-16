package main

import (
	"context"
	"net/http"

	"github.com/darksasori/graphql/db"
	"github.com/darksasori/graphql/repository"
	"github.com/darksasori/graphql/schema"
	"github.com/darksasori/graphql/service"
	"github.com/graph-gophers/graphql-go/relay"

	"log"
)

func main() {
	conn, err := db.Connect(context.TODO())
	if err != nil {
		panic(err)
	}
	connDB := conn.Database("test")

	user := service.NewUser(repository.NewUser(connDB))
	s := schema.New(user)

	log.Println("Listen on :8080")
	http.Handle("/graphql", &relay.Handler{s})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
