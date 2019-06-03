
dev:
	go build -tags dev -o bin/dev cmd/test.go
	./bin/dev
