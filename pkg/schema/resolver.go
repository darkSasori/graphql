package schema

import (
	"github.com/darksasori/graphql/pkg/service"
	graphql "github.com/graph-gophers/graphql-go"
)

type Resolver struct {
	user *service.User
}

func New(user *service.User) *graphql.Schema {
	r := &Resolver{
		user,
	}
	return graphql.MustParseSchema(schemaSctring, r)
}
