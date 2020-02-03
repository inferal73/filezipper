VERSION = `git describe --abbrev=0 --tags`

.PHONY: build
build:
	go build -o ./dist/filezipper -ldflags "-X main.Version=${VERSION}" ./cmd/filezipper

.PHONY: test
test:
	go test -v -race -timeout 30s ./...