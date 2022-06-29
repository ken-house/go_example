.PHONY: build
build:
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/server main.go

.PHONY: clean
clean:
        rm -f ./bin/server