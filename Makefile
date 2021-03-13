APP := "vulcan"
CURRENT_VERSION := $(shell git branch --show-current | cut -d '/' -f2)

test:
	@go test -cover ./...

#To use go lint you must install `go get -u github.com/golangci/golangci-lint/cmd/golangci-lint`
lint:
	@golangci-lint run -E golint -E bodyclose ./...

clean:
	@rm -rf bin/

build: clean
	@go build \
		-ldflags "-X 'github.com/caioeverest/vulcan/infra/config.Version=${CURRENT_VERSION}'" \
		-o dist/${APP} main.go

install: build
	cp dist/${APP} ${HOME}/.local/bin

# To use mockery you must install `go get -u go get github.com/vektra/mockery/v2/.../`
update-mocks:
	@mockery --all --inpackage --case=underscore

make-release-local:
	goreleaser --snapshot --skip-publish --rm-dist
