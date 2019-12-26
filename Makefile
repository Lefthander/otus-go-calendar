PROJECT?=github.com/lefthander/otus-go-calendar

GOOS?=darwin
GOARCH?=amd64

RELEASE?=0.0.0
BUILD_NUMBER := git-$(shell git rev-list HEAD --count)
COMMIT := git-$(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 go build \
		-ldflags "-s -w -X ${PROJECT}/internal/version.Version=${RELEASE} \
		-X ${PROJECT}/internal/version.BuildNumber=${BUILD_NUMBER} \
		-X ${PROJECT}/internal/version.CommitHash=${COMMIT} \
		-X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
		-o bin/otus-go-calendar ${PROJECT}/cmd/otus-go-calendar