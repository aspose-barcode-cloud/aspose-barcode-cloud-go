.PHONY: all
all: release

.PHONY: init
init:
	./scripts/init.sh

.PHONY: format
format:
	./scripts/fix_api_error.sh
	./scripts/format.sh
	./scripts/docs_format.sh

.PHONY: test
test:
	./scripts/test.sh

.PHONY: build
build:
	./scripts/build.sh

.PHONY: lint
lint:
	./scripts/lint.sh

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
after-gen: init format insert-examples clean-gomod

.PHONY: insert-examples
insert-examples:
	./scripts/insert-examples.bash

.PHONY: ci
ci: build test
