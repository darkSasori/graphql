package main

import (
	"net/http"

	"log"

	"github.com/darksasori/graphql"
)

func main() {
	http.HandleFunc("/graphql", graphql.Graphql)

	log.Println("Listen on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
