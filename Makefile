
.PHONY: all
all: format lint test

.PHONY: format
format:
	./scripts/fmt.sh

.PHONY: lint
lint:
	./scripts/vet.sh

.PHONY: test
test:
	./scripts/test.sh

.PHONY: build
build: format lint
	./scripts/build.sh

.PHONY: clean
clean:
	./scripts/clean.sh
