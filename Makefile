# Current project parameters
CURRENT_PROJECT=github.com/amarin/payload/

# Go parameters
GOCMD=GO111MODULE=on go
DEPCMD=$(GOCMD) mod
VERSION=${VER}
GOTEST=$(GOCMD) test

.PHONY: all help deps build clean

help: ## display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lint: ## lint the files
	@golangci-lint run

test: ## unittest code with
	@export GOPATH=$GOPATH:`pwd`
	@$(GOTEST) -v ./...

tidy: ## Fix dependencies records in go.mod
	@$(DEPCMD) tidy

godoc: ## Run godoc locally
	@echo "Starting local godoc http server at http://127.0.0.1:6060/pkg/${CURRENT_PROJECT}"
	@godoc -http=:6060
