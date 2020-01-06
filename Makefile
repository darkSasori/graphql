test-unit:
	go test -v --tags test_unit -race ./...

lint:
	golint -set_exit_status ./...

server:
	go build -tags dev -o bin/server cmd/server.go
	./bin/server

deploy:
	gcloud functions deploy graphql \
		--entry-point Graphql \
		--runtime go111 \
		--trigger-http \
		--env-vars-file envs.yml

mongodb:
	docker run --rm --name "mongodb" -d -p 27017:27017 mongo
