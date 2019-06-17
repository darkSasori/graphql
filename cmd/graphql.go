package main

import (
	"context"
	"html/template"
	"net/http"

	"github.com/darksasori/graphql/db"
	"github.com/darksasori/graphql/repository"
	"github.com/darksasori/graphql/schema"
	"github.com/darksasori/graphql/service"
	"github.com/graph-gophers/graphql-go/relay"

	"log"
)

type playgroundData struct {
	PlaygroundVersion    string
	Endpoint             string
	SubscriptionEndpoint string
	SetTitle             bool
}

func main() {
	conn, err := db.Connect(context.TODO())
	if err != nil {
		panic(err)
	}
	connDB := conn.Database("test")

	user := service.NewUser(repository.NewUser(connDB))
	s := schema.New(user)
	handler := &relay.Handler{Schema: s}

	t := template.New("Playground")
	t, err = t.Parse(playgroundTemplate)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-apollo-tracing")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		switch r.Method {
		case http.MethodPost:
			handler.ServeHTTP(w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		case http.MethodGet:
			d := playgroundData{
				PlaygroundVersion: "1.5.2",
				Endpoint:          r.URL.Path,
				SetTitle:          true,
			}
			err = t.ExecuteTemplate(w, "index", d)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("Listen on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

const playgroundTemplate = `
{{ define "index" }}
<!--
The request to this GraphQL server provided the header "Accept: text/html"
and as a result has been presented Playground - an in-browser IDE for
exploring GraphQL.
If you wish to receive JSON, provide the header "Accept: application/json" or
add "&raw" to the end of the URL within a browser.
-->
<!DOCTYPE html>
<html>
<head>
  <meta charset=utf-8/>
  <meta name="viewport" content="user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, minimal-ui">
  <title>GraphQL Playground</title>
  <link rel="stylesheet" href="//cdn.jsdelivr.net/npm/graphql-playground-react/build/static/css/index.css" />
  <link rel="shortcut icon" href="//cdn.jsdelivr.net/npm/graphql-playground-react/build/favicon.png" />
  <script src="//cdn.jsdelivr.net/npm/graphql-playground-react/build/static/js/middleware.js"></script>
</head>
<body>
  <div id="root">
    <style>
      body {
        background-color: rgb(23, 42, 58);
        font-family: Open Sans, sans-serif;
        height: 90vh;
      }
      #root {
        height: 100%;
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
      }
      .loading {
        font-size: 32px;
        font-weight: 200;
        color: rgba(255, 255, 255, .6);
        margin-left: 20px;
      }
      img {
        width: 78px;
        height: 78px;
      }
      .title {
        font-weight: 400;
      }
    </style>
    <img src='//cdn.jsdelivr.net/npm/graphql-playground-react/build/logo.png' alt=''>
    <div class="loading"> Loading
      <span class="title">GraphQL Playground</span>
    </div>
  </div>
  <script>window.addEventListener('load', function (event) {
      GraphQLPlayground.init(document.getElementById('root'), {
        // options as 'endpoint' belong here
        endpoint: {{ .Endpoint }},
        setTitle: {{ .SetTitle }}
      })
    })</script>
</body>
</html>
{{ end }}
`
