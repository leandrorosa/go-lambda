build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/hello.go hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go