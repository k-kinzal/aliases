USER     := k-kinzal
REPO     := aliases
GIT_TAG  := $(shell git tag --points-at HEAD)
GIT_HASH := $(shell git rev-parse HEAD)
VERSION  := $(shell if [ -n "$(GIT_TAG)" ]; then echo "$(GIT_TAG)"; else echo "$(GIT_HASH)"; fi)

DIST_DIR := $(shell if [ -n "$(GOOS)$(GOARCH)" ]; then echo "./dist/$(GOOS)-$(GOARCH)"; else echo "./dist"; fi)

.PHONY: build
build:
	go build -ldflags "-X github.com/$(USER)/$(REPO)/pkg/version.version=$(VERSION)" -o $(DIST_DIR)/aliases .

.PHONY: cross-build
cross-build:
	@make build GOOS=linux GOARCH=amd64
	@make build GOOS=darwin GOARCH=amd64

.PHONY: test
test:
	go test ./... -v


.PHONY: clean
clean:
	@rm -rf $(DIST_DIR)
