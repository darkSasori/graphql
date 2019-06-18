
dev:
	go build -tags dev -o bin/dev cmd/test.go
	./bin/dev

test-unit:
	go test -v --tags test_unit -race github.com/darksasori/graphql/pkg/...

lint:
	golint -set_exit_status ./...


graphql:
	go build -tags dev -o bin/graphql cmd/graphql.go
	./bin/graphql

test:
	curl -XPOST -d '{"query": "{ listUsers { id, displayname } }"}' localhost:8080/graphql |jq
