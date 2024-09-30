BINARY=exify

install-deps:
	go install github.com/abice/go-enum@latest

generate:
	go generate ./...

test: 
	go test -cover ./...

build:
	go build -o bin/${BINARY} ./cmd/server/main.go

build-cli:
	go build -o bin/${BINARY}-cli ./cmd/cli/main.go

unittest:
	go test -short -v ./...

clean:
	if [ -d bin ] ; then rm -rf bin ; fi

format:
	go fmt ./...

analyze:
	go vet ./...


lint-prepare:
	@echo "Installing golangci-lint" 
	go install honnef.co/go/tools/cmd/staticcheck@latest

lint:
	staticcheck ./...

install-dependencies:
	go install github.com/pressly/goose/v3/cmd/goose@v3.13.4 && go mod download 


run-dev:
	go run ./cmd/server/main.go

MIGRATIONNAME = ""
MIGRATIONTYPE = "sql"

create-migration:
	goose -dir ./migrations/mysql create $(MIGRATIONNAME) $(MIGRATIONTYPE)


release-staging:
	standard-version -p beta

release-prod:
	standard-version