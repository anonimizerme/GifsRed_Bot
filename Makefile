.PHONY: build clean deploy deploy_simple_bot

build:
#	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
#	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/gifsred_bot src/lambdas/gifsred_bot/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

deploy_simple_bot: clean build
	sls deploy --verbose -f gifsred_bot