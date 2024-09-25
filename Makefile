.PHONY: help mocks linter test swag.gen.init swag.merge.init swag.gen swag.merge swag.init swag
.DEFAULT: help # Running Make will run the help target

help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


mocks: ## Generate mocks
	mockery --all --recursive --keeptree

linter: ## Run linter
	golangci-lint run

test: ## run tests
	go test -v -cover ./...

swag.gen.init: ## install swagger generator
	@echo :: getting generator
	go get -d github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	go mod tidy

swag.merge.init: ## install swagger merge
	@echo :: init merge api yaml docs
	rm -f $(which yq)
	GO111MODULE=on go get -d github.com/mikefarah/yq/v4
	go install github.com/mikefarah/yq/v4

swag.gen: ## generate swagger
	./api/swag-gen.sh

swag.merge: ## merge swagger
	@echo :: merging api yaml docs
	./api/swag-merge.sh

swag.init: swag.gen.init swag.merge.init
swag: swag.gen swag.merge
