.PHONY: build clean deploy

build:
	dep ensure -v

	env GOOS=linux go build -ldflags="-s -w" -o bin/getAllTodos src/functions/getAllTodos/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/getTodoWithID src/functions/getTodoWithID/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/insertTodo src/functions/insertTodo/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/deleteTodoWithID src/functions/deleteTodoWithID/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

local: 
	serverless sam export --output ./template.yml

deploy-local:
	sam local start-api

deploy-aws: clean build
	sls deploy --verbose
