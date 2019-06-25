test-unit:
	go test -v --tags test_unit -race ./...

lint:
	golint -set_exit_status ./...

server:
	go build -tags dev -o bin/server cmd/server.go
	./bin/server

test:
	curl -XPOST -d '{"query": "{ listUsers { id, displayname } }"}' localhost:8080/graphql |jq

deploy:
	gcloud alpha functions deploy graphql \
		--entry-point Graphql \
		--runtime go111 \
		--trigger-http \
		--env-vars-file=env
