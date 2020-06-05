SRC=./aspose_barcode_cloud
TEST=./test

.PHONY: all
all: format test

.PHONY: format
format:
	./scripts/fmt.sh

.PHONY: test
test:
	./scripts/test.sh

.PHONY: build
build: format
	./scripts/build.sh

.PHONY: clean
clean:
	./scripts/clean.sh
