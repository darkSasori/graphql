package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/darksasori/graphql/pkg/db"
	"github.com/darksasori/graphql/pkg/model"
	"github.com/darksasori/graphql/pkg/repository"
	"github.com/darksasori/graphql/pkg/schema"
	"github.com/darksasori/graphql/pkg/service"
	"github.com/darksasori/graphql/pkg/utils"
	"github.com/graph-gophers/graphql-go/relay"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var handler *relay.Handler
var conf *oauth2.Config
var tmplError *template.Template
var tmplPlayground *template.Template
var userService *service.User

func init() {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	conn, err := db.Connect(ctx)
	if err != nil {
		panic(err)
	}

	db := utils.GetEnvDefault("MONGODB_NAME", "gcloud")
	connDB := conn.Database(db)

	userService = service.NewUser(repository.NewUser(connDB))
	s := schema.New(userService)
	handler = &relay.Handler{Schema: s}

	tmplError, err = template.New("error").Parse(errorHTML)
	if err != nil {
		panic(err)
	}

	tmplPlayground, err = template.New("playground").Parse(playgroundHTML)
	if err != nil {
		panic(err)
	}

	conf = &oauth2.Config{
		ClientID:     utils.GetEnvDefault("CLIENT_ID", ""),
		ClientSecret: utils.GetEnvDefault("CLIENT_SECRET", ""),
		RedirectURL:  utils.GetEnvDefault("REDIRECT_URL", ""),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
}

func saveUser(tokenType, accessToken string) *model.User {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v1/userinfo?alt=json", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", tokenType, accessToken))

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
		return nil
	}

	decoder := json.NewDecoder(resp.Body)

	var user struct {
		Username string `json:"email"`
		Name     string `json:"name"`
		Image    string `json:"picture"`
	}
	if err := decoder.Decode(&user); err != nil {
		panic(err)
	}

	u := model.NewUser(user.Username, user.Name, user.Image)
	err = userService.Save(context.TODO(), u)
	if err != nil {
		panic(err)
	}

	return u
}

// Graphql handle request
func Graphql(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-apollo-tracing")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	switch r.Method {
	case http.MethodPost:
		auth := strings.Split(r.Header.Get("Authorization"), " ")
		user := saveUser(auth[0], auth[1])
		if user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(w, r.WithContext(utils.NewUserContext(r.Context(), user)))
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		queries := r.URL.Query()
		if _, ok := queries["state"]; ok {
			ctx := context.Background()
			code := queries["code"]
			token, err := conf.Exchange(ctx, code[0])
			if err != nil {
				tmplError.Execute(w, nil)
				fmt.Printf("Error: %s\n", err)
			} else {
				saveUser(token.Type(), token.AccessToken)

				url := fmt.Sprintf("/graphql?token=%s&type=%s", token.AccessToken, token.Type())
				http.Redirect(w, r, url, 302)
			}
		} else if token, ok := queries["token"]; ok {
			accessToken := token[0]
			tokenType := queries["type"][0]
			if saveUser(tokenType, accessToken) == nil {
				tmplError.Execute(w, nil)
				return
			}

			tmplPlayground.Execute(w, struct{ Type, AccessToken string }{tokenType, accessToken})
		} else {
			url := conf.AuthCodeURL("state")
			http.Redirect(w, r, url, 302)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

const playgroundHTML = `
<!DOCTYPE html>
<html>
	<head>
		<title>Graphql</title>
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
					headers: {
						'Authorization': '{{ .Type }} {{ .AccessToken }}'
					},
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

const errorHTML = `
<!DOCTYPE html>
<html>
	<head>
		<title>Graphql</title>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<p>Error: {{ .Error }}</p>
		<a href='/graphql'>Login</a>
	</body>
</html>
`
