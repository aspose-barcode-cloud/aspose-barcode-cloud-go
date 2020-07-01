.PHONY: all
all: format test

.PHONY: format
format:
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

.PHONY: tidy
tidy:
	./scripts/tidy.sh

.PHONY: update_packages
update_packages:
	./scripts/update_packages.sh

.PHONY: update
update: update_packages tidy

.PHONY: release
release: format vet lint update_packages tidy build test

.PHONY: ci
ci: build test
