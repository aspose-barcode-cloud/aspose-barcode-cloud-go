
.PHONY: all
all: format vet test tidy

.PHONY: format
format:
	./scripts/fmt.sh

.PHONY: vet
vet:
	./scripts/vet.sh

.PHONY: test
test:
	./scripts/test.sh

.PHONY: build
build: format vet
	./scripts/build.sh

.PHONY: lint
lint: format
	./scripts/lint.sh

.PHONY: tidy
tidy:
	./scripts/tidy.sh
