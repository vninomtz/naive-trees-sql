build:
	go build -o bin/tserver ./cmd/server/main.go

run: build 
	./bin/tserver
