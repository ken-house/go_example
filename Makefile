.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -ldflags "-s -w" -tags=jsoniter -o ./bin/go-example main.go

.PHONY: push
push:
	docker build . -t xudengtang/go-example:latest
	docker push xudengtang/go-example:latest

.PHONY: clean
clean:
	rm -f ./bin/go-example
	docker image rm xudengtang/go-example:latest