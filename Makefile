
dev:
	go build -tags dev -o bin/dev cmd/test.go
	./bin/dev

test-unit:
	go test -v --tags test_unit -race ./...

lint:
	golint -set_exit_status ./...
