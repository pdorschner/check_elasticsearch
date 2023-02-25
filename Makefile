.PHONY: test coverage lint vet

build:
	go build -ldflags="-X 'main.version=new_value'"
goreleaser:
	goreleaser release --snapshot --rm-dis
lint:
	go fmt $(go list ./... | grep -v /vendor/)
vet:
	go vet $(go list ./... | grep -v /vendor/)
test:
	go test -v -cover ./...
coverage:
	go test -v -cover -coverprofile=coverage.out ./... &&\
	go tool cover -html=coverage.out -o coverage.html
