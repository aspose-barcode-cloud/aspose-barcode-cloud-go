.PHONY: all
all: release

.PHONY: init
init:
	go get -v -t -d ./...
	go install golang.org/x/tools/cmd/goimports@v0.16.0
	go install honnef.co/go/tools/cmd/staticcheck@v0.3.3

.PHONY: format
format:
	./scripts/fix_api_error.sh
	$$(go env GOPATH)/bin/goimports -w .
	./scripts/docs_format.sh

.PHONY: test
test:
	./scripts/test.sh

.PHONY: build
build:
	./scripts/build.sh

.PHONY: lint
lint:
	./scripts/vet.sh
	$$(go env GOPATH)/bin/staticcheck ./...

.PHONY: update_packages
update_packages:
	./scripts/update_packages.sh

.PHONY: clean-gosum
clean-gosum: clean-gomod
	./scripts/clean-go.sum.sh

.PHONY: clean-gomod
clean-gomod:
	./scripts/clean-go.mod.sh
	./scripts/tidy.sh

.PHONY: update
update: update_packages clean-gomod

.PHONY: release
release: format lint update_packages clean-gomod build test

.PHONY: after-gen
after-gen: init format clean-gomod

.PHONY: ci
ci: build test
