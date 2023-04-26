.PHONY: all
all: release

.PHONY: init
init:
	go get -v -t -d ./...

.PHONY: format
format:
	./scripts/fix_api_error.sh
	./scripts/fmt.sh
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

.PHONY: update_packages
update_packages:
	./scripts/update_packages.sh

.PHONY: tidy
tidy:
	./scripts/tidy.sh

.PHONY: clean-gosum
clean-gosum: clean-gomod
	./scripts/clean-go.sum.sh

.PHONY: clean-gomod
clean-gomod:
	./scripts/clean-go.mod.sh

.PHONY: update
update: update_packages tidy clean-gosum

.PHONY: release
release: format lint update_packages tidy build test clean-gosum

.PHONY: after-gen
after-gen: init format clean-gosum

.PHONY: ci
ci: build test
