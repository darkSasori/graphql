package graphql

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/darksasori/graphql/pkg/db"
	"github.com/darksasori/graphql/pkg/repository"
	"github.com/darksasori/graphql/pkg/schema"
	"github.com/darksasori/graphql/pkg/service"
	"github.com/darksasori/graphql/pkg/utils"
	"github.com/graph-gophers/graphql-go/relay"
)

var handler *relay.Handler

func init() {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()
	conn, err := db.Connect(ctx)
	if err != nil {
		panic(err)
	}

	db := utils.GetEnvDefault("MONGODB_NAME", "gcloud")
	connDB := conn.Database(db)

	user := service.NewUser(repository.NewUser(connDB))
	s := schema.New(user)
	handler = &relay.Handler{Schema: s}
}

func Graphql(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-apollo-tracing")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	switch r.Method {
	case http.MethodPost:
		handler.ServeHTTP(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		io.WriteString(w, playgroundHtml)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

const playgroundHtml = `
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/graphql", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`
