.PHONY: all
all: release

.PHONY: format
format:
	./scripts/fix_api_error.sh
	./scripts/fmt.sh
	./scripts/docs_format.sh

.PHONY: vet
vet:
	./scripts/vet.sh

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

.PHONY: tidy
tidy:
	./scripts/tidy.sh

.PHONY: clean-gosum
clean-gosum:
	./scripts/clean-go.sum.sh

.PHONY: clean-gomod
clean-gomod:
	./scripts/clean-go.mod.sh

.PHONY: update
update: update_packages tidy clean-gomod clean-gosum

.PHONY: release
release: format vet lint update_packages tidy build test clean-gomod clean-gosum

.PHONY: after-gen
after-gen: format clean-gomod clean-gosum

.PHONY: ci
ci: build test
